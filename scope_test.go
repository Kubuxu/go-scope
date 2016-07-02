package scope

import (
	"testing"
)

func TestPanic(t *testing.T) {
	closeA := SubScope("A")
	closeB := SubScope("B")
	assertPanic(t, closeA)
	closeB()
	closeA()
}

func TestScopeName(t *testing.T) {
	defer SubScope("ext")()
	if Current().Name() != "ext" {
		t.Fatalf("Expected scope name: %s, got: %s", "ext", Current().Name())
	}
	closeInt := SubScope("int")
	if Current().Name() != "ext.int" {
		t.Fatalf("Expected scope name: %s, got: %s", "ext.int", Current().Name())
	}
	closeInt()
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
