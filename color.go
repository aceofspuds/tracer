package tracer

// Color is a tuple with an r, g, b coordinate.
func Color(r, g, b float64) Tuple {
	return Tuple([]float64{r, g, b})
}

// // r returns the first element in a tuple.
// func (t Tuple) r() float64 { return t[0] }

// // g returns the second element in a tuple.
// func (t Tuple) g() float64 { return t[1] }

// // b returns the third element in a tuple.
// func (t Tuple) b() float64 { return t[2] }

// Product computes the hadamard product of two colors.
func (t Tuple) Product(t1 Tuple) Tuple {
	out := Tuple(make([]float64, len(t)))
	for i := range t {
		out[i] = t[i] * t1[i]
	}
	return out
}
