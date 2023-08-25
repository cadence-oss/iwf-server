package integ

import (
	"context"
	"github.com/indeedeng/iwf/integ/workflow/locking"
	"github.com/indeedeng/iwf/service/common/ptr"
	"strconv"
	"testing"
	"time"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/service"
	"github.com/stretchr/testify/assert"
)

func TestLockingWorkflowTemporal(t *testing.T) {
	if !*temporalIntegTest {
		t.Skip()
	}
	for i := 0; i < *repeatIntegTest; i++ {
		doTestLockingWorkflow(t, service.BackendTypeTemporal, nil)
		smallWaitForFastTest()
	}
}

func TestLockingWorkflowTemporalContinueAsNew(t *testing.T) {
	if !*temporalIntegTest {
		t.Skip()
	}
	for i := 0; i < *repeatIntegTest; i++ {
		doTestLockingWorkflow(t, service.BackendTypeTemporal, minimumContinueAsNewConfig())
		smallWaitForFastTest()
	}
}

func TestLockingWorkflowCadence(t *testing.T) {
	if !*cadenceIntegTest {
		t.Skip()
	}
	for i := 0; i < *repeatIntegTest; i++ {
		doTestLockingWorkflow(t, service.BackendTypeCadence, nil)
		smallWaitForFastTest()
	}
}

func TestLockingWorkflowCadenceContinueAsNew(t *testing.T) {
	if !*cadenceIntegTest {
		t.Skip()
	}
	for i := 0; i < *repeatIntegTest; i++ {
		doTestLockingWorkflow(t, service.BackendTypeCadence, minimumContinueAsNewConfig())
		smallWaitForFastTest()
	}
}

func doTestLockingWorkflow(t *testing.T, backendType service.BackendType, config *iwfidl.WorkflowConfig) {
	// start test workflow server
	wfHandler := locking.NewHandler()
	closeFunc1 := startWorkflowWorker(wfHandler)
	defer closeFunc1()

	_, closeFunc2 := startIwfServiceByConfig(IwfServiceTestConfig{
		BackendType:                      backendType,
		DisableFailAtMemoIncompatibility: true,
	})
	defer closeFunc2()

	// start a workflow
	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: "http://localhost:" + testIwfServerPort,
			},
		},
	})
	wfId := locking.WorkflowType + strconv.Itoa(int(time.Now().UnixNano()))

	req := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background())
	startReq := iwfidl.WorkflowStartRequest{
		WorkflowId:             wfId,
		IwfWorkflowType:        locking.WorkflowType,
		WorkflowTimeoutSeconds: 100,
		IwfWorkerUrl:           "http://localhost:" + testWorkflowServerPort,
		StartStateId:           ptr.Any(locking.State1),
		WorkflowStartOptions: &iwfidl.WorkflowStartOptions{
			WorkflowConfigOverride: config,
		},
	}
	_, httpResp, err := req.WorkflowStartRequest(startReq).Execute()
	panicAtHttpError(err, httpResp)

	assertions := assert.New(t)

	for i := 0; i < 10; i++ {
		allSearchAttributes := []iwfidl.SearchAttributeKeyAndType{
			{
				Key:       iwfidl.PtrString(locking.TestSearchAttributeKeywordKey),
				ValueType: iwfidl.KEYWORD.Ptr(),
			},
			{
				Key:       iwfidl.PtrString(locking.TestSearchAttributeIntKey),
				ValueType: iwfidl.INT.Ptr(),
			},
		}
		time.Sleep(time.Second * 1)
		reqRpc := apiClient.DefaultApi.ApiV1WorkflowRpcPost(context.Background())
		rpcResp, httpResp, err := reqRpc.WorkflowRpcRequest(iwfidl.WorkflowRpcRequest{
			WorkflowId: wfId,
			RpcName:    locking.RPCName,
			Input:      locking.TestValue,
			SearchAttributesLoadingPolicy: &iwfidl.PersistenceLoadingPolicy{
				PersistenceLoadingType: iwfidl.PARTIAL_WITH_EXCLUSIVE_LOCK.Ptr(),
				PartialLoadingKeys: []string{
					locking.TestSearchAttributeKeywordKey,
				},
				LockingKeys: []string{
					locking.TestSearchAttributeIntKey,
				},
			},
			DataAttributesLoadingPolicy: &iwfidl.PersistenceLoadingPolicy{
				PersistenceLoadingType: iwfidl.PARTIAL_WITH_EXCLUSIVE_LOCK.Ptr(),
				PartialLoadingKeys: []string{
					locking.TestDataObjectKey2,
				},
				LockingKeys: []string{
					locking.TestDataObjectKey1,
				},
			},
			TimeoutSeconds:   iwfidl.PtrInt32(2),
			SearchAttributes: allSearchAttributes,
		}).Execute()
		if err != nil || httpResp.StatusCode != 200 {
			panicAtHttpError(err, httpResp)
		}
		assertions.Equal(rpcResp.Output, locking.TestValue)
	}

	req2 := apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(context.Background())
	resp2, httpResp, err := req2.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId: wfId,
	}).Execute()
	panicAtHttpError(err, httpResp)

	history, _ := wfHandler.GetTestResult()
	assertions.Equalf(map[string]int64{
		"S1_start":  1,
		"S1_decide": 1,
		"S2_start":  10,
		"S2_decide": 10,
	}, history, "locking.test fail, %v", history)

	assertions.Equal(iwfidl.COMPLETED, resp2.GetWorkflowStatus())
	assertions.Equal(10, len(resp2.GetResults()))

	reqSearch := apiClient.DefaultApi.ApiV1WorkflowSearchattributesGetPost(context.Background())
	searchResult2, httpResp, err := reqSearch.WorkflowGetSearchAttributesRequest(iwfidl.WorkflowGetSearchAttributesRequest{
		WorkflowId: wfId,
		Keys: []iwfidl.SearchAttributeKeyAndType{
			{
				Key:       iwfidl.PtrString(locking.TestSearchAttributeIntKey),
				ValueType: ptr.Any(iwfidl.INT),
			},
		},
	}).Execute()
	panicAtHttpError(err, httpResp)

	expectedSearchAttributeInt := iwfidl.SearchAttribute{
		Key:          iwfidl.PtrString(locking.TestSearchAttributeIntKey),
		ValueType:    ptr.Any(iwfidl.INT),
		IntegerValue: iwfidl.PtrInt64(10),
	}
	assertions.Equal([]iwfidl.SearchAttribute{expectedSearchAttributeInt}, searchResult2.GetSearchAttributes())

	reqQry := apiClient.DefaultApi.ApiV1WorkflowDataobjectsGetPost(context.Background())
	// force to test compatibility of memo
	useMemo := false
	if backendType == service.BackendTypeTemporal {
		useMemo = true
	}
	queryResult1, httpResp, err := reqQry.WorkflowGetDataObjectsRequest(iwfidl.WorkflowGetDataObjectsRequest{
		WorkflowId: wfId,
		Keys: []string{
			locking.TestDataObjectKey1,
		},
		UseMemoForDataAttributes: ptr.Any(useMemo),
	}).Execute()
	panicAtHttpError(err, httpResp)

	expected1 := []iwfidl.KeyValue{
		{
			Key: iwfidl.PtrString(locking.TestDataObjectKey1),
			Value: &iwfidl.EncodedObject{
				Encoding: iwfidl.PtrString("json"),
				Data:     iwfidl.PtrString("10"),
			},
		},
	}
	assertions.ElementsMatch(expected1, queryResult1.GetObjects())
}
