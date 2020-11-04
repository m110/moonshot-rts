package systems_test

import (
	"testing"

	"github.com/m110/moonshot-rts/internal/engine"
	"github.com/m110/moonshot-rts/internal/systems"
	"github.com/stretchr/testify/require"
)

func TestEntityList_Remove(t *testing.T) {
	list := systems.EntityList{}

	a := engine.NewBaseEntity()
	b := engine.NewBaseEntity()
	c := engine.NewBaseEntity()
	d := engine.NewBaseEntity()

	list.Add(a)
	list.Add(b)
	list.Add(c)
	list.Add(d)

	require.Equal(t, []engine.Entity{a, b, c, d}, list.All())

	list.Remove(c)
	require.Equal(t, []engine.Entity{a, b, d}, list.All())

	list.Remove(a)
	require.Equal(t, []engine.Entity{b, d}, list.All())

	list.Remove(d)
	require.Equal(t, []engine.Entity{b}, list.All())

	list.Remove(b)
	require.Equal(t, []engine.Entity{}, list.All())

	// Idempotency
	list.Remove(b)
	require.Equal(t, []engine.Entity{}, list.All())
}
