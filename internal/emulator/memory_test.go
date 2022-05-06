package emulator

import "testing"

func TestLength(t *testing.T) {
	want := memorySize
	got := len(memory{}.storage)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSetGet(t *testing.T) {
	var want uint8

	m := memory{}
	index := 23
	m.Set(index, 42)
	want = 42
	got := m.Get(index)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMemoryLowerBound(t *testing.T) {
	defer func() { recover() }()

	memory{}.Get(-1)

	// Unreachable if `Get` panics as intended
	t.Errorf("did not panic")
}

func TestMemoryUpperBound(t *testing.T) {
	defer func() { recover() }()

	memory{}.Get(memorySize + 1)

	// Unreachable if `Get` panics as intended
	t.Errorf("did not panic")
}