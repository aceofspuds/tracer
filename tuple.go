package tracer

import "math"

// Tuple represents a coordinate with some number of values.
type Tuple []float64

// Equal returns true if each coordinate of the tuple is within some
// epsilon of its counterpart.
func (t Tuple) Equal(t1 Tuple, e float64) bool {
	if len(t) != len(t1) {
		return false
	}
	for i := range t {
		if !eq(t[i], t1[i], e) {
			return false
		}
	}
	return true
}

// Add adds the coordinates of two tuples.
func (t Tuple) Add(t1 Tuple) Tuple {
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] + t1[i]
	}
	return out
}

// Sub subtracts the coordinates of two tuples.
func (t Tuple) Sub(t1 Tuple) Tuple {
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] - t1[i]
	}
	return out
}

// Multiply multiplies the coordinates of a tuple by a scalar value.
func (t Tuple) Multiply(s float64) Tuple {
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] * s
	}
	return out
}

// Divide divides the coordinates of a tuple by a scalar value.
func (t Tuple) Divide(s float64) Tuple {
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] / s
	}

	return out
}

// Negate negates a tuple.
func (t Tuple) Negate() Tuple {
	return t.Multiply(-1)
}

// Magnitude returns the magnitude of a tuple.
func (t Tuple) Magnitude() float64 {
	sum := 0.
	for i := range t {
		sum += t[i] * t[i]
	}
	return math.Sqrt(sum)
}

// Normalize normalizes a vector.
func (t Tuple) Normalize() Tuple {
	// find magnitude
	m := t.Magnitude()
	if m == 0. {
		m = 1.
	}

	// normalize
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] / m
	}

	return out
}

// IsZero returns true if all values a tuple are zero.
func (t Tuple) IsZero() bool {
	for _, v := range t {
		if v != 0. {
			return false
		}
	}

	return true
}

func eq(f1, f2, e float64) bool {
	return math.Abs(f1-f2) < e
}
