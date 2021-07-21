package tuple

import (
	"math"
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
		{Point(3, -2, 5).Add(Vector(-2, 3, 1)), Point(1, 1, 6)},
		{Point(3, 2, 1).Sub(Point(5, 6, 7)), Vector(-2, -4, -6)},
		{Vector(3, 2, 1).Sub(Vector(5, 6, 7)), Vector(-2, -4, -6)},
		{Tuple{1, -2, 3, -4}.Negate(), Tuple{-1, 2, -3, 4}},
		{Tuple{1, -2, 3, -4}.Multiply(3.5), Tuple{3.5, -7, 10.5, -14}},
		{Tuple{1, -2, 3, -4}.Multiply(0.5), Tuple{0.5, -1, 1.5, -2}},
		{Tuple{1, -2, 3, -4}.Divide(2), Tuple{0.5, -1, 1.5, -2}},
		{Vector(1, 0, 0).Multiply(Vector(1, 2, 3).Magnitude()), Vector(math.Sqrt(14), 0, 0)},
		{Vector(1, 0, 0).Normalize(), Vector(1, 0, 0)},
		{Vector(4, 0, 0).Normalize(), Vector(1, 0, 0)},
		{Vector(1, 0, 0).Multiply(Vector(1, 2, 3).Dot(Vector(2, 3, 4))), Vector(20, 0, 0)},
		{Vector(1, 2, 3).Cross(Vector(2, 3, 4)), Vector(-1, 2, -1)},
		{Vector(2, 3, 4).Cross(Vector(1, 2, 3)), Vector(1, -2, 1)},
		{Vector(2, 3, 4).Cross(Vector(1, 2, 3)).Add(Vector(1, 2, 3).Cross(Vector(2, 3, 4))), Vector(0, 0, 0)},
	}

	for i, td := range tds {
		if !Equal(td.input, td.expected, epsilon) {
			t.Errorf("test %d failed: expected tuples to be equal, return %v %v", i, td.input, td.expected)
		}
	}
}
