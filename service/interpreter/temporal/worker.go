package temporal

import (
	"github.com/indeedeng/iwf/service/common/config"
	"github.com/indeedeng/iwf/service/interpreter"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

type InterpreterWorker struct {
	temporalClient client.Client
	worker         worker.Worker
	taskQueue      string
}

func NewInterpreterWorker(config config.Config, temporalClient client.Client, taskQueue string) *InterpreterWorker {
	interpreter.SetSharedConfig(config)
	return &InterpreterWorker{
		temporalClient: temporalClient,
		taskQueue:      taskQueue,
	}
}

func (iw *InterpreterWorker) Close() {
	iw.temporalClient.Close()
	iw.worker.Stop()
}

func (iw *InterpreterWorker) Start() {
	config := interpreter.GetSharedConfig()
	options := worker.Options{
		MaxConcurrentActivityTaskPollers: 10,
		MaxConcurrentWorkflowTaskPollers: 10,
	}
	if config.Interpreter.Temporal != nil && config.Interpreter.Temporal.WorkerOptions != nil {
		options = *config.Interpreter.Temporal.WorkerOptions
	}
	iw.worker = worker.New(iw.temporalClient, iw.taskQueue, options)
	worker.EnableVerboseLogging(config.Interpreter.VerboseDebug)

	iw.worker.RegisterWorkflow(Interpreter)
	iw.worker.RegisterActivity(interpreter.StateStart)
	iw.worker.RegisterActivity(interpreter.StateDecide)
	iw.worker.RegisterActivity(interpreter.DumpWorkflowInternal)

	err := iw.worker.Start()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
