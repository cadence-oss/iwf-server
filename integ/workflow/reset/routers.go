package reset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/helpers"
	"github.com/indeedeng/iwf/service"
	"log"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

/**
* This test workflow has 2 states, using REST controller to implement the workflow directly.
* State1:
*       - No WaitUntil
*       - Execute moves to State2
* State2:
* 		- No WaitUntil
*       - Execute loops through state2 5 times, then gracefully completes the workflow.
* This test is used for testing reset by state id and state execution id without WaitUntil
 */
const (
	WorkflowType = "reset"
	State1       = "S1"
	State2       = "S2"
)

type handler struct {
	invokeHistory sync.Map
}

func NewHandler() *handler {
	return &handler{
		invokeHistory: sync.Map{},
	}
}

// ApiV1WorkflowStartPost - for a workflow
func (h *handler) ApiV1WorkflowStateStart(c *gin.Context, t *testing.T) {
	helpers.FailTestWithErrorMessage("No start call is expected.", t)
}

func (h *handler) ApiV1WorkflowStateDecide(c *gin.Context, t *testing.T) {
	log.Println("start of decide")
	var req iwfidl.WorkflowStateDecideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("received state decide request, ", req)
	context := req.GetContext()
	if context.GetAttempt() <= 0 || context.GetFirstAttemptTimestamp() <= 0 {
		helpers.FailTestWithErrorMessage("attempt and firstAttemptTimestamp should be greater than zero", t)
	}

	if req.GetWorkflowType() == WorkflowType {
		if value, ok := h.invokeHistory.Load(req.GetWorkflowStateId() + "_decide"); ok {
			h.invokeHistory.Store(req.GetWorkflowStateId()+"_decide", value.(int64)+1)
		} else {
			h.invokeHistory.Store(req.GetWorkflowStateId()+"_decide", int64(1))
		}

		if req.GetWorkflowStateId() == State1 {
			// go to S2
			c.JSON(http.StatusOK, iwfidl.WorkflowStateDecideResponse{
				StateDecision: &iwfidl.StateDecision{
					NextStates: []iwfidl.StateMovement{
						{
							StateId:    State2,
							StateInput: req.StateInput,
							StateOptions: &iwfidl.WorkflowStateOptions{
								//Skipping wait until for 1st execution of state2
								SkipWaitUntil: iwfidl.PtrBool(true),
							},
						},
					},
				},
			})
			return
		} else if req.GetWorkflowStateId() == State2 {
			input := req.GetStateInput()
			i, _ := strconv.Atoi(input.GetData())
			if i < 5 {
				updatedInput := &iwfidl.EncodedObject{
					Encoding: iwfidl.PtrString("json"),
					Data:     iwfidl.PtrString(fmt.Sprintf("%v", i+1)),
				}
				c.JSON(http.StatusOK, iwfidl.WorkflowStateDecideResponse{
					StateDecision: &iwfidl.StateDecision{
						NextStates: []iwfidl.StateMovement{
							{
								StateId:    State2,
								StateInput: updatedInput,
								StateOptions: &iwfidl.WorkflowStateOptions{
									//Skipping wait until for all executions of state2 after the 1st execution.
									SkipWaitUntil: iwfidl.PtrBool(true),
								},
							},
						},
					},
				})
				return
			} else {
				// go to complete
				c.JSON(http.StatusOK, iwfidl.WorkflowStateDecideResponse{
					StateDecision: &iwfidl.StateDecision{
						NextStates: []iwfidl.StateMovement{
							{
								StateId:    service.GracefulCompletingWorkflowStateId,
								StateInput: req.StateInput,
							},
						},
					},
				})
				return
			}
		}
	}

	c.JSON(http.StatusBadRequest, struct{}{})
}

func (h *handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	invokeHistory := make(map[string]int64)
	h.invokeHistory.Range(func(key, value interface{}) bool {
		invokeHistory[key.(string)] = value.(int64)
		return true
	})
	return invokeHistory, nil
}
