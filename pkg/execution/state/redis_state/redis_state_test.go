package redis_state

import (
	"context"
	"crypto/rand"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/inngest/inngest/pkg/event"
	"github.com/inngest/inngest/pkg/execution/state"
	"github.com/inngest/inngest/pkg/execution/state/testharness"
	"github.com/oklog/ulid/v2"
	"github.com/rueian/rueidis"
	"github.com/stretchr/testify/require"
)

func TestStateHarness(t *testing.T) {
	r := miniredis.RunT(t)
	sm, err := New(
		context.Background(),
		WithKeyPrefix("{test}:"),
		WithFunctionLoader(testharness.FunctionLoader()),
		WithConnectOpts(rueidis.ClientOption{
			InitAddress:  []string{r.Addr()},
			DisableCache: true,
		}),
	)
	require.NoError(t, err)

	create := func() (state.Manager, func()) {
		return sm, func() {
			r.FlushAll()
		}
	}

	testharness.CheckState(t, create)
}

func BenchmarkNew(b *testing.B) {
	r := miniredis.RunT(b)
	sm, err := New(
		context.Background(),
		WithConnectOpts(rueidis.ClientOption{
			InitAddress:  []string{r.Addr()},
			DisableCache: true,
		}),
	)
	require.NoError(b, err)

	id := state.Identifier{
		WorkflowID: uuid.New(),
	}
	init := state.Input{
		Identifier: id,
		EventData: event.Event{
			Name: "test-event",
			Data: map[string]any{
				"title": "They don't think it be like it is, but it do",
				"data": map[string]any{
					"float": 3.14132,
				},
			},
			User: map[string]any{
				"external_id": "1",
			},
			Version: "1985-01-01",
		}.Map(),
	}

	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		init.Identifier.RunID = ulid.MustNew(ulid.Now(), rand.Reader)
		_, err := sm.New(ctx, init)
		require.NoError(b, err)
	}

}
