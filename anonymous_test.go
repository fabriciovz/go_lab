package go_lab

import (
	"testing"
)

func TestAnonymous(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		MainChain10()
	})
	t.Run("", func(t *testing.T) {
		MainChain11()
	})
	t.Run("", func(t *testing.T) {
		MainChain12()
	})
}
