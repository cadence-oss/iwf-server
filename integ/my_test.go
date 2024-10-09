package integ

import (
	"github.com/indeedeng/iwf/service"
	"testing"
	"time"
)

// remove the leading underscore when using it
func _TestNothinButJustRunningTheServiceTemporalWorkerForDebug(t *testing.T) {
	startIwfServiceWithClient(service.BackendTypeTemporal)
	time.Sleep(time.Hour)
}