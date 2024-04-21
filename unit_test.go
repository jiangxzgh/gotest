package gotest_test

import (
	"gotest"
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = a + b

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d): expected %d, actual %d", a, b, expected, actual)
	}
}
