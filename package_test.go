package code_test

import (
	"testing"

	"code"
)

func TestIncrement(t *testing.T) {
	if code.Increment(1) != 2 {
		t.Fail()
	}
}
