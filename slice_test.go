package from_test

import (
	"maps"
	"testing"

	"github.com/martindrlik/from"
)

func TestSlice(t *testing.T) {
	t.Run("slice index as key", func(t *testing.T) {
		key := func(i int, _ string) int {
			return i
		}
		input := []string{"foo", "bar"}
		actual := from.Slice(input, key, from.NoValue)
		expect := map[int]struct{}{
			0: {},
			1: {},
		}
		if !maps.Equal(actual, expect) {
			t.Errorf("expected %+v, got %+v", expect, actual)
		}
	})
	t.Run("slice value as key", func(t *testing.T) {
		key := func(_ int, s string) string {
			return s
		}
		value := func(int, string) struct{} {
			return struct{}{}
		}
		input := []string{"foo", "bar"}
		actual := from.Slice(input, key, value)
		expect := map[string]struct{}{
			"foo": {},
			"bar": {},
		}
		if !maps.Equal(actual, expect) {
			t.Errorf("expected %+v, got %+v", expect, actual)
		}
	})
	t.Run("slice value as key, index as value", func(t *testing.T) {
		key := func(_ int, s string) string {
			return s
		}
		value := func(i int, _ string) int {
			return i
		}
		input := []string{"foo", "bar"}
		actual := from.Slice(input, key, value)
		expect := map[string]int{
			"foo": 0,
			"bar": 1,
		}
		if !maps.Equal(actual, expect) {
			t.Errorf("expected %+v, got %+v", expect, actual)
		}
	})
}
