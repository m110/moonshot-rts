package engine_test

import (
	"testing"

	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/stretchr/testify/require"
)

type intEvent struct {
	value int
}

type stringEvent struct {
	value string
}

type subscriber struct {
	ints    []int
	strings []string
}

func (s *subscriber) HandleEvent(e engine.Event) {
	switch event := e.(type) {
	case intEvent:
		s.ints = append(s.ints, event.value)
	case stringEvent:
		s.strings = append(s.strings, event.value)
	default:
		panic("received unknown event")
	}
}

func TestEventBus(t *testing.T) {
	eb := engine.NewEventBus()

	s1 := &subscriber{}
	s2 := &subscriber{}

	eb.Subscribe(intEvent{}, s1)
	eb.Subscribe(stringEvent{}, s1)

	eb.Subscribe(intEvent{}, s2)

	eb.Publish(intEvent{1})
	eb.Publish(intEvent{2})
	eb.Publish(intEvent{3})

	eb.Publish(stringEvent{"a"})
	eb.Publish(stringEvent{"b"})
	eb.Publish(stringEvent{"c"})

	require.Equal(t, []int{1, 2, 3}, s1.ints)
	require.Equal(t, []string{"a", "b", "c"}, s1.strings)

	require.Equal(t, []int{1, 2, 3}, s2.ints)
	require.Nil(t, s2.strings)
}
