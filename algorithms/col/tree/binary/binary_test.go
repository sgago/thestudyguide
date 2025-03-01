package binary

import "testing"

func TestTotal(t *testing.T) {
	if total := Total(0); total != 1 {
		t.Errorf("Expected 1, got %d", total)
	}

	if total := Total(1); total != 3 {
		t.Errorf("Expected 3, got %d", total)
	}

	if total := Total(2); total != 7 {
		t.Errorf("Expected 7, got %d", total)
	}

	if total := Total(3); total != 15 {
		t.Errorf("Expected 15, got %d", total)
	}

	if total := Total(4); total != 31 {
		t.Errorf("Expected 31, got %d", total)
	}
}

func TestParentIdx(t *testing.T) {
	if parent := ParentIdx(1); parent != 0 {
		t.Errorf("Expected 0, got %d", parent)
	}

	if parent := ParentIdx(2); parent != 0 {
		t.Errorf("Expected 0, got %d", parent)
	}

	if parent := ParentIdx(3); parent != 1 {
		t.Errorf("Expected 1, got %d", parent)
	}

	if parent := ParentIdx(4); parent != 1 {
		t.Errorf("Expected 1, got %d", parent)
	}

	if parent := ParentIdx(5); parent != 2 {
		t.Errorf("Expected 2, got %d", parent)
	}

	if parent := ParentIdx(6); parent != 2 {
		t.Errorf("Expected 2, got %d", parent)
	}

	if parent := ParentIdx(7); parent != 3 {
		t.Errorf("Expected 3, got %d", parent)
	}

	if parent := ParentIdx(8); parent != 3 {
		t.Errorf("Expected 3, got %d", parent)
	}
}

func TestLeftChildIdx(t *testing.T) {
	if left := LeftChildIdx(0); left != 1 {
		t.Errorf("Expected 1, got %d", left)
	}

	if left := LeftChildIdx(1); left != 3 {
		t.Errorf("Expected 3, got %d", left)
	}

	if left := LeftChildIdx(2); left != 5 {
		t.Errorf("Expected 5, got %d", left)
	}

	if left := LeftChildIdx(3); left != 7 {
		t.Errorf("Expected 7, got %d", left)
	}

	if left := LeftChildIdx(4); left != 9 {
		t.Errorf("Expected 9, got %d", left)
	}
}

func TestRightChildIdx(t *testing.T) {
	if right := RightChildIdx(0); right != 2 {
		t.Errorf("Expected 2, got %d", right)
	}

	if right := RightChildIdx(1); right != 4 {
		t.Errorf("Expected 4, got %d", right)
	}

	if right := RightChildIdx(2); right != 6 {
		t.Errorf("Expected 6, got %d", right)
	}

	if right := RightChildIdx(3); right != 8 {
		t.Errorf("Expected 8, got %d", right)
	}

	if right := RightChildIdx(4); right != 10 {
		t.Errorf("Expected 10, got %d", right)
	}
}
