package function

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/gosimple/slug"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/inngest/inngestctl/inngest"
	"github.com/inngest/inngestctl/inngest/state"
)

// Function represents a step function which is triggered whenever an event
// is received or on a schedule.  In essence, it contains:
//
// - Triggers, which represent when a function is invoked
//
// - Steps, which represent the individual steps of actions that the function calls.
//
// A function may be simple (ie. only having a single step) or complex (ie. many
// steps).  Simple functions are easy:  run the single step's action.  Complex functions
// represent steps as a DAG, with edges between the trigger and each step.
type Function struct {
	// Name is the descriptive name for the function
	Name string `json:"name"`

	// ID is the immutable random ID for the function.
	ID string `json:"id"`

	// Trigger represnets the trigger for the function.
	Triggers []Trigger `json:"triggers"`

	// Idempotency allows the specification of an idempotency key by templating event
	// data, eg:
	//
	//  `{{ event.data.order_id }}`.
	//
	// When specified, a function will run at most once per 24 hours for the given unique
	// key.
	Idempotency *string `json:"idempotency,omitempty"`

	// Throttle allows specifying custom throttling for the function.
	Throttle *inngest.Throttle `json:"throttle,omitempty"`

	// Actions represents the actions to take for this function.  If empty, this assumes
	// that we have a single action specified in the current directory using
	Steps map[string]Step `json:"steps,omitempty"`
}

// Step represents a single unit of code (action) which runs as part of a step function, in a DAG.
type Step struct {
	Name    string                 `json:"name"`
	Runtime inngest.RuntimeWrapper `json:"runtime"`
}

// New returns a new, empty function with a randomly generated ID.
func New() (*Function, error) {
	id, err := randomID()
	if err != nil {
		return nil, err
	}
	return &Function{
		ID:    id,
		Steps: map[string]Step{},
	}, nil
}

func (f Function) Slug() string {
	return slug.Make(f.Name)
}

// Validate returns an error if the function definition is invalid.
func (f Function) Validate() error {
	var err error
	if f.ID == "" {
		err = multierror.Append(err, fmt.Errorf("A function ID is required"))
	}
	if f.Name == "" {
		err = multierror.Append(err, fmt.Errorf("A function name is required"))
	}
	if len(f.Triggers) == 0 {
		err = multierror.Append(err, fmt.Errorf("At least one trigger is required"))
	}
	for _, t := range f.Triggers {
		if terr := t.Validate(); terr != nil {
			err = multierror.Append(err, terr)
		}
	}
	return err
}

// Workflow produces the workflow.cue definition for a function.  Our executor
// runs a "workflow", which is a DAG of the function steps.  Its a subset of
// the function used purely for execution.
func (f Function) Workflow(ctx context.Context) (*inngest.Workflow, error) {
	w := inngest.Workflow{
		Name:     f.Name,
		ID:       f.ID,
		Triggers: make([]inngest.Trigger, len(f.Triggers)),
	}

	for n, t := range f.Triggers {
		if t.EventTrigger != nil {
			w.Triggers[n].EventTrigger = &inngest.EventTrigger{
				Event:      t.EventTrigger.Event,
				Expression: t.EventTrigger.Expression,
			}
			continue
		}
		if t.CronTrigger != nil {
			w.Triggers[n].CronTrigger = &inngest.CronTrigger{
				Cron: t.CronTrigger.Cron,
			}
			continue
		}
		return nil, fmt.Errorf("unknown trigger type")
	}

	if f.Throttle != nil {
		w.Throttle = f.Throttle
	}

	if f.Idempotency != nil {
		w.Throttle = &inngest.Throttle{
			Key:    f.Idempotency,
			Count:  1,
			Period: "24h",
		}
	}

	// This has references to actions.  Create the actions then reference them
	// from the workflow.
	actions, edges, err := f.Actions(ctx)
	if err != nil {
		return nil, err
	}

	for n, a := range actions {
		w.Actions = append(w.Actions, inngest.Action{
			ClientID: uint(n) + 1, // 0 is the trigger; use 1 offset
			Name:     a.Name,
			DSN:      a.DSN,
		})
	}

	w.Edges = edges

	// TODO: When supporting > 1 action, validate the edges.

	return &w, nil
}

// Actions produces configuration for each step of the function.  Each config
// file specifies how to run the code.
func (f Function) Actions(ctx context.Context) ([]inngest.ActionVersion, []inngest.Edge, error) {
	// This has no defined actions, which means its an implicit
	// single action invocation.  We assume that a Dockerfile
	// exists in the project root, and that we can build the
	// image which contains all of the code necessary to run
	// the function.
	if len(f.Steps) == 0 {
		return f.defaultAction(ctx)
	}

	// XXX: We're redefining how we define "edges" within steps, and so right now
	// only support 1 edge.
	if len(f.Steps) > 1 {
		return nil, nil, fmt.Errorf("Function definitions currently only support one step")
	}

	// TODO: With > 1 step, convert each step into a tree, with nodes pointing from the trigger
	// to each step in a stable manner.

	var (
		a   inngest.ActionVersion
		err error
	)

	for _, step := range f.Steps {
		if a, err = f.action(ctx, step, 0); err != nil {
			return nil, nil, err
		}
	}

	edges := []inngest.Edge{{
		Outgoing: "trigger",
		Incoming: 1,
	}}
	return []inngest.ActionVersion{a}, edges, nil
}

func (f Function) action(ctx context.Context, s Step, n int) (inngest.ActionVersion, error) {
	suffix := "test"
	if state.IsProd() {
		suffix = "prod"
	}

	id := fmt.Sprintf("%s-step-%d-%s", f.ID, n, suffix)
	if prefix, err := state.AccountIdentifier(ctx); err == nil {
		id = fmt.Sprintf("%s/%s", prefix, id)
	}

	a := inngest.ActionVersion{
		Name:    s.Name,
		DSN:     id,
		Runtime: s.Runtime,
	}
	if s.Runtime.RuntimeType() != "http" {
		// Non-HTTP actions can read secrets;  http actions are external APIs and so
		// don't need secret access.
		a.Scopes = []string{"secret:read:*"}
	}
	return a, nil
}

func (f Function) defaultAction(ctx context.Context) ([]inngest.ActionVersion, []inngest.Edge, error) {
	id := f.ID + "-action"

	if prefix, err := state.AccountIdentifier(ctx); err == nil {
		id = fmt.Sprintf("%s/%s", prefix, id)
	}

	a := inngest.ActionVersion{
		Name: f.Name,
		DSN:  id,
		// This is a custom action, so allow reading any secret.
		Scopes: []string{"secret:read:*"},
		Runtime: inngest.RuntimeWrapper{
			Runtime: inngest.RuntimeDocker{
				Image: id,
			},
		},
	}
	edges := []inngest.Edge{{
		Outgoing: "trigger",
		Incoming: 1,
	}}
	return []inngest.ActionVersion{a}, edges, nil
}

func randomID() (string, error) {
	// Generate a 6 character long hex string.  This is the suffix to
	// our DSN prefix, which decreases the chance of collosion by 1/16,777,216.
	// This makes the total chance of collisions from an _empty_ keyspace
	// 1 in 3,435,034,312,704 (we'll ignore birthday problems).
	byt := make([]byte, 3)
	if _, err := rand.Read(byt); err != nil {
		return "", fmt.Errorf("error generating ID: %w", err)
	}
	petname.NonDeterministicMode()
	return fmt.Sprintf("%s-%s", petname.Generate(2, "-"), hex.EncodeToString(byt)), nil
}
