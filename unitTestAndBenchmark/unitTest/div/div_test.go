package div

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {
	res, err := Div(1, 2)
	if err != nil {
		t.Errorf("error (%+v)", err)
		return
	}
	fmt.Println(res)
}

func TestAll(t *testing.T) {
	_, err := Div(1, 2)
	if err != nil {
		t.Errorf("error (%+v)", err)
		return
	}
	_, err = Div(1, 0)
	if err == nil {
		t.Errorf("error (%+v)", err)
		return
	}
}

func TestDivs(t *testing.T) {
	tables := []struct {
		x float64
		y float64
		n float64
	}{
		{1, 1, 1},
		{1, 2, 0.5},
		{4, 2, 2},
		{10, 2, 5},
	}
	for _, table := range tables {
		res, err := Div(table.x, table.y)
		if err != nil {
			t.Errorf("error (%+v)", err)
			return
		}
		if res != table.n {
			t.Errorf("Div of (%f/%f) was incorrect, got: %f, want: %f.", table.x, table.y, res, table.n)
		}
	}
}
