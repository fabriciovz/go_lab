package go_lab

import (
	"testing"
)

func TestWaitGroup(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		MainChain30()
	})
	t.Run("", func(t *testing.T) {
		MainChain31()
	})
	t.Run("", func(t *testing.T) {
		MainChain32()
	})
}
