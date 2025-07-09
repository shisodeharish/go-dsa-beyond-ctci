package dynamicarray

import "testing"

func TestDynamicArrayExtras(t *testing.T) {
	arr := NewDynamicArrayExtras()
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)

	val, err := arr.Pop(1)
	if err != nil || val != 2 {
		t.Errorf("Expected 2, got %d, err: %v", val, err)
	}
	if v, _ := arr.Get(1); v != 3 {
		t.Errorf("Expected 3 after pop, got %d", v)
	}

	if !arr.Contains(1) || arr.Contains(10) {
		t.Errorf("Contains failed")
	}

	err = arr.Insert(1, 99)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	if v, _ := arr.Get(1); v != 99 {
		t.Errorf("Expected 99, got %d", v)
	}

	index := arr.Remove(99)
	if index != 1 {
		t.Errorf("Expected index 1 for removed 99, got %d", index)
	}
	if v, _ := arr.Get(1); v == 99 {
		t.Errorf("Remove failed, 99 still present")
	}
}