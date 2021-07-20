package tuple

import (
	"testing"
)

const epsilon = 0.00000001

func TestTuple(t *testing.T) {
	p := Tuple{4.3, -4.2, 3.1, 1.0}

	if p.x != 4.3 {
		t.Errorf("expected 4.3, return %.f", p.x)
	}
	if p.y != -4.2 {
		t.Errorf("expected -4.2, return %.f", p.y)
	}
	if p.z != 3.1 {
		t.Errorf("expected 3.1, return %.f", p.z)
	}
	if p.w != 1.0 {
		t.Errorf("expected 1.0, return %.f", p.w)
	}
}

func TestDefinitions(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	tds := []test{
		{Point(4, -4, 3), Tuple{4, -4, 3, 1}},
		{Vector(4, -4, 3), Tuple{4, -4, 3, 0}},
		{Add(Point(3, -2, 5), Vector(-2, 3, 1)), Point(1, 1, 6)},
		{Sub(Point(3, 2, 1), Point(5, 6, 7)), Vector(-2, -4, -6)},
	}

	for _, td := range tds {
		if !Equal(td.input, td.expected, epsilon) {
			t.Errorf("expected tuples to be equal, return %v %v", td.input, td.expected)
		}
	}
}
