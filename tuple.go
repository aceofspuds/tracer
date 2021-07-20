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
func Add(t1, t2 Tuple) Tuple {
	return Tuple{t1.x + t2.x, t1.y + t2.y, t1.z + t2.z, t1.w + t2.w}
}

// Sub subtracts the coordinates of two tuples.
func Sub(t1, t2 Tuple) Tuple {
	return Tuple{t1.x - t2.x, t1.y - t2.y, t1.z - t2.z, t1.w - t2.w}
}

func eq(f1, f2, e float64) bool {
	return math.Abs(f1-f2) < e
}
