package main

import (
	"encoding/json"
	"testing"

	"github.com/taskcluster/taskcluster-client-go/queue"
)

// Badly formatted json payload should result in *json.SyntaxError error in task.validatePayload()
func TestBadPayloadValidate(t *testing.T) {

	// replace task update channels to use a dummy updater, in order to consume messages
	taskStatusUpdate, taskStatusUpdateErr = func() (request chan<- TaskStatusUpdate, err <-chan error) {
		r := make(chan TaskStatusUpdate)
		e := make(chan error)
		go func() {
			for {
				<-r
				e <- nil
			}
		}()
		return r, e
	}()

	badPayload := json.RawMessage(`bad payload, not even json`)
	task := TaskRun{Definition: queue.TaskDefinitionResponse{Payload: badPayload}}
	err := task.validatePayload()
	if err == nil {
		t.Fatalf("Bad task payload should not have passed validation")
	}
	switch err.(type) {
	case *json.SyntaxError:
		t.Log("Received *json.SyntaxError as expected - all ok!")
	default:
		t.Errorf("Bad task payload should have retured a *json.SyntaxError error, but actually returned a %T error. The unexpected %T error was:\n%s", err, err, err)
	}
}