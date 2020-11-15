package utils

import (
	"testing"
)

func TestValidVersion(t *testing.T) {
	t.Run("valid situation", func(t *testing.T) {
		got := ValidVersion("1.0.0")
		want := true
		if got != want {
			t.Errorf("want %t, got %t", want, got)
		}
	})
}
