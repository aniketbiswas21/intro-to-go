package utils

import (
	"testing"
)

// The test should always accept the same argument, t, of type *testing.T)
func TestAdd(t *testing.T) {
  // the value you expect to get from the tested function
  expected := 6
  actual := Add(1, 2, 3)
  if actual != expected {
    t.Errorf("Add was incorrect! Expected: %d, Actual: %d", expected, actual)
  }
}