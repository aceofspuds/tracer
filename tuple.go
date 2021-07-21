package tuple

import "math"

// Tuple represents an (x, y, z, w) coordinate, where w is 0 if
// the it is point a and 1 if it is a vector.
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

// Point creates a point tuple from x, y, and z coordinates.
func Point(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 1}
}

// Vector creates a vector tuple from x, y, and z coordinates.
func Vector(x, y, z float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: 0}
}

// Equal returns true if each coordinate of the tuple is within some
// epsilon of its counterpart.
func Equal(t1, t2 Tuple, e float64) bool {
	return eq(t1.x, t2.x, e) && eq(t1.y, t2.y, e) && eq(t1.z, t2.z, e) && eq(t1.w, t2.w, e)
}

// Add adds the coordinates of two tuples.
func (t Tuple) Add(t1 Tuple) Tuple {
	return Tuple{t.x + t1.x, t.y + t1.y, t.z + t1.z, t.w + t1.w}
}

// Sub subtracts the coordinates of two tuples.
func (t Tuple) Sub(t1 Tuple) Tuple {
	return Tuple{t.x - t1.x, t.y - t1.y, t.z - t1.z, t.w - t1.w}
}

// Multiply multiplies the coordinates of a tuple by a scalar value.
func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.x * s, t.y * s, t.z * s, t.w * s}
}

// Divide divides the coordinates of a tuple by a scalar value.
func (t Tuple) Divide(s float64) Tuple {
	return Tuple{t.x / s, t.y / s, t.z / s, t.w / s}
}

// Negate negates a tuple.
func (t Tuple) Negate() Tuple {
	return t.Multiply(-1)
}

// Magnitude returns the magnitude of a tuple.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt((t.x * t.x) + (t.y * t.y) + (t.z * t.z) + (t.w * t.w))
}

// Normalize normalizes a vector.
func (t Tuple) Normalize() Tuple {
	return t.Divide(t.Magnitude())
}

// Dot comptes the dot product of two tuples.
func (t Tuple) Dot(t1 Tuple) float64 {
	return (t.x * t1.x) + (t.y * t1.y) + (t.z * t1.z) + (t.w * t1.w)
}

// Cross comptes the cross product of two tuples.
func (t Tuple) Cross(t1 Tuple) Tuple {
	return Vector(t.y*t1.z-t.z*t1.y, t.z*t1.x-t.x*t1.z, t.x*t1.y-t.y*t1.x)
}

func eq(f1, f2, e float64) bool {
	return math.Abs(f1-f2) < e
}
