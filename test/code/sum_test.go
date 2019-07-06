package code

import "testing"

func TestSum(t *testing.T) {
	total, err := Div(1, 1)
	if err != nil {
		t.Errorf("error %s", err.Error())
		return
	}
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSums(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{4, 2, 2},
		{8, 2, 4},
	}

	for _, table := range tables {
		total, err := Div(table.x, table.y)
		if err != nil {
			t.Errorf("error %s", err.Error())
			return
		}
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
