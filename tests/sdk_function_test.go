package main

import (
	"testing"
	"time"

	"github.com/inngest/inngest/inngest"
	"github.com/inngest/inngest/pkg/execution/driver"
	"github.com/inngest/inngest/pkg/function"
	"github.com/inngest/inngestgo"
)

func TestSDKFunctions(t *testing.T) {
	// 1. Assert that a function is registered with the name of "sdk-function-test"
	// 1. Assert that there's an invocation with no steps.
	evt := inngestgo.Event{
		Name: "tests/function.test",
		Data: map[string]any{
			"test": true,
		},
		User: map[string]any{},
	}

	fnID := "test-suite-sdk-function-test"
	test := &Test{
		Name: "SDK Functions",
		Description: `
			This test asserts that functions work across SDKs.

			In order for functions to work, the SDK must:
			- Allow functions to be introspected via the GET handler
			- Register functions correctly with the server
			- Handle incoming event triggers
			- Respond with the correct, expected data
		`,
		Function: function.Function{
			ID:   fnID,
			Name: "SDK Function Test",
			Triggers: []function.Trigger{
				{
					EventTrigger: &function.EventTrigger{
						Event: "tests/function.test",
					},
				},
			},
			Steps: map[string]function.Step{
				"step": {
					ID:   "step",
					Name: "step",
					Runtime: &inngest.RuntimeWrapper{
						Runtime: &inngest.RuntimeHTTP{
							URL: stepURL(fnID, "step"),
						},
					},
				},
			},
		},
		EventTrigger: evt,
	}

	test.SetAssertions(
		// All executor requests should have this event.
		test.SetRequestEvent(evt),
		// And the executor should start its requests with this context.
		test.SetRequestContext(SDKCtx{
			FnID:   fnID,
			StepID: "step",
			Stack: driver.FunctionStack{
				Current: 0,
			},
		}),
		test.SendTrigger(),
		test.ExpectRequest("Initial request", "step", time.Second),
		test.ExpectJSONResponse(200, map[string]any{"name": "tests/function.test", "body": "ok"}),
	)

	run(t, test)
}