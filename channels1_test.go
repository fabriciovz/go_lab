package go_lab

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		MainChain0()
	})
	t.Run("", func(t *testing.T) {
		MainChain1()
	})
	t.Run("", func(t *testing.T) {
		MainChain2()
	})
	t.Run("", func(t *testing.T) {
		MainChain3()
	})
	t.Run("", func(t *testing.T) {
		MainChain5()
	})
}
