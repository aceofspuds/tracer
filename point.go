package tracer

// Point is a tuple with an x, y, z coordinate and a forth value w=1.
func Point(x, y, z float64) Tuple {
	return Tuple([]float64{x, y, z, 1})
}

// Vector is a tuple with an x, y, z coordinate and a forth value w=0.
func Vector(x, y, z float64) Tuple {
	return Tuple([]float64{x, y, z, 0})
}

// x returns the first element in a tuple.
func (t Tuple) x() float64 { return t[0] }

// y returns the second element in a tuple.
func (t Tuple) y() float64 { return t[1] }

// z returns the third element in a tuple.
func (t Tuple) z() float64 { return t[2] }

// w returns the fourth element in a tuple.
func (t Tuple) w() float64 { return t[3] }

// Dot computes the dot product of two tuples.
func (t Tuple) Dot(t1 Tuple) float64 {
	sum := 0.
	for i := range t {
		sum += t[i] * t1[i]
	}
	return sum
}

// Cross comptes the cross product of two tuples.
func (t Tuple) Cross(t1 Tuple) Tuple {
	return Vector(
		t.y()*t1.z()-t.z()*t1.y(),
		t.z()*t1.x()-t.x()*t1.z(),
		t.x()*t1.y()-t.y()*t1.x())
}
