package integ

import (
	"context"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/persistence"
	"github.com/indeedeng/iwf/integ/workflow/signal"
	"github.com/indeedeng/iwf/service"
	"github.com/indeedeng/iwf/service/common/ptr"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestSetSearchAttributes(t *testing.T) {
	if !*temporalIntegTest {
		t.Skip()
	}
	assertions := assert.New(t)

	// start test workflow server
	wfHandler := signal.NewHandler()
	closeFunc1 := startWorkflowWorker(wfHandler, t)
	defer closeFunc1()

	_, closeFunc2 := startIwfServiceWithClient(service.BackendTypeTemporal)
	defer closeFunc2()

	wfId := signal.WorkflowType + strconv.Itoa(int(time.Now().UnixNano()))

	// start a workflow
	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: "http://localhost:" + testIwfServerPort,
			},
		},
	})
	req := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background())
	_, httpResp, err := req.WorkflowStartRequest(iwfidl.WorkflowStartRequest{
		WorkflowId:             wfId,
		IwfWorkflowType:        signal.WorkflowType,
		WorkflowTimeoutSeconds: 10,
		IwfWorkerUrl:           "http://localhost:" + testWorkflowServerPort,
		StartStateId:           ptr.Any(signal.State1),
	}).Execute()
	failTestAtHttpError(err, httpResp, t)

	assertions.Equal(httpResp.StatusCode, http.StatusOK)

	var signalVals []iwfidl.SearchAttribute
	signalVals = append(signalVals, iwfidl.SearchAttribute{
		Key:          iwfidl.PtrString(persistence.TestSearchAttributeIntKey),
		ValueType:    ptr.Any(iwfidl.INT),
		IntegerValue: iwfidl.PtrInt64(persistence.TestSearchAttributeIntValue1),
	},
		iwfidl.SearchAttribute{
			Key:         iwfidl.PtrString(persistence.TestSearchAttributeKeywordKey),
			ValueType:   ptr.Any(iwfidl.KEYWORD),
			StringValue: iwfidl.PtrString(persistence.TestSearchAttributeKeywordValue1),
		},
		iwfidl.SearchAttribute{
			Key:              iwfidl.PtrString(persistence.TestSearchAttributeKeywordArrayKey),
			ValueType:        ptr.Any(iwfidl.KEYWORD_ARRAY),
			StringArrayValue: []string{persistence.TestSearchAttributeKeywordValue2, persistence.TestSearchAttributeKeywordValue1},
		})

	setReq := apiClient.DefaultApi.ApiV1WorkflowSearchattributesSetPost(context.Background())
	httpResp2, err := setReq.WorkflowSetSearchAttributesRequest(iwfidl.WorkflowSetSearchAttributesRequest{
		WorkflowId:       wfId,
		SearchAttributes: signalVals,
	}).Execute()

	failTestAtHttpError(err, httpResp2, t)

	time.Sleep(time.Second)

	getReq := apiClient.DefaultApi.ApiV1WorkflowSearchattributesGetPost(context.Background())
	searchResult, httpRespGet, err := getReq.WorkflowGetSearchAttributesRequest(iwfidl.WorkflowGetSearchAttributesRequest{
		WorkflowId: wfId,
		Keys: []iwfidl.SearchAttributeKeyAndType{
			{
				Key:       iwfidl.PtrString(persistence.TestSearchAttributeIntKey),
				ValueType: ptr.Any(iwfidl.INT),
			},
			{
				Key:       iwfidl.PtrString(persistence.TestSearchAttributeKeywordKey),
				ValueType: ptr.Any(iwfidl.KEYWORD),
			},
			{
				Key:       iwfidl.PtrString(persistence.TestSearchAttributeKeywordArrayKey),
				ValueType: ptr.Any(iwfidl.KEYWORD_ARRAY),
			},
		}}).Execute()
	failTestAtHttpError(err, httpRespGet, t)

	assertions.ElementsMatch(signalVals, searchResult.SearchAttributes)

	// Terminate the workflow once tests completed
	stopReq := apiClient.DefaultApi.ApiV1WorkflowStopPost(context.Background())
	_, err = stopReq.WorkflowStopRequest(iwfidl.WorkflowStopRequest{
		WorkflowId: wfId,
		StopType:   iwfidl.TERMINATE.Ptr(),
	}).Execute()
}
