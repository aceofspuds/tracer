package tracer

import (
	"math"
	"testing"
)

func Test4x4Matrix(t *testing.T) {
	m := Matrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})

	type test struct {
		row      int
		col      int
		expected float64
	}

	tds := []test{
		{0, 0, 1},
		{0, 3, 4},
		{1, 0, 5.5},
		{1, 2, 7.5},
		{2, 2, 11},
		{3, 0, 13.5},
		{3, 2, 15.5},
	}

	for i, td := range tds {
		output := m[td.row][td.col]
		if output != td.expected {
			t.Errorf("test %d failed: expected %f, returned %f", i, td.expected, output)
		}
	}
}

func Test2x2Matrix(t *testing.T) {
	m := Matrix([][]float64{
		{-3, 5},
		{1, -2},
	})

	type test struct {
		row      int
		col      int
		expected float64
	}

	tds := []test{
		{0, 0, -3},
		{0, 1, 5},
		{1, 0, 1},
		{1, 1, -2},
	}

	for i, td := range tds {
		output := m[td.row][td.col]
		if output != td.expected {
			t.Errorf("test %d failed: expected %f, returned %f", i, td.expected, output)
		}
	}
}

func TestMatrixEqual(t *testing.T) {
	m1 := Matrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})

	m2 := Matrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})

	if !m1.Equal(m2, epsilon) {
		t.Error("expected equal, returned false")
		return
	}

	m2[3][3] = 0

	if m1.Equal(m2, epsilon) {
		t.Error("expected not equal, returned true")
		return
	}
}

func TestMatrixMultiply(t *testing.T) {
	m1 := Matrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})

	m2 := Matrix([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})

	output := m1.Multiply(m2)
	expected := Matrix([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}

	output = output.Multiply(IdentityMatrix(4))
	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}
	output = IdentityMatrix(4).Multiply(output)
	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}
}

func TestMatrixMultiplyT(t *testing.T) {
	m1 := Matrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})

	t1 := Tuple([]float64{1, 2, 3, 1})

	outputTuple := m1.MultiplyT(t1)
	expectedTuple := Tuple([]float64{18, 24, 33, 1})

	if !outputTuple.Equal(expectedTuple, epsilon) {
		t.Errorf("expected %v, returned %v", expectedTuple, outputTuple)
		return
	}

	outputTuple = IdentityMatrix(4).MultiplyT(outputTuple)
	if !outputTuple.Equal(expectedTuple, epsilon) {
		t.Errorf("expected %v, returned %v", expectedTuple, outputTuple)
		return
	}

	m2 := Matrix([][]float64{
		{1},
		{2},
		{3},
		{1},
	})

	outputMatrix := m1.Multiply(m2)
	expectedMatrix := Matrix([][]float64{
		{18},
		{24},
		{33},
		{1},
	})

	if !outputMatrix.Equal(expectedMatrix, epsilon) {
		t.Errorf("expected %v, returned %v", expectedMatrix, outputMatrix)
		return
	}
	outputMatrix = IdentityMatrix(4).Multiply(outputMatrix)
	if !outputMatrix.Equal(expectedMatrix, epsilon) {
		t.Errorf("expected %v, returned %v", expectedMatrix, outputMatrix)
		return
	}
}

func TestMatrixTranspose(t *testing.T) {
	m := Matrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})

	output := m.Transpose()
	expected := Matrix([][]float64{
		{1, 5, 9, 5},
		{2, 6, 8, 4},
		{3, 7, 7, 3},
		{4, 8, 6, 2},
	})

	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}

	m = IdentityMatrix(4)
	output = m.Transpose()
	if !output.Equal(m, epsilon) {
		t.Errorf("expected %v, returned %v", m, output)
		return
	}
}

func TestSubMatrix(t *testing.T) {
	m := Matrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})

	output := m.SubMatrix(0, 2)
	expected := Matrix([][]float64{
		{5, 6, 8},
		{9, 8, 6},
		{5, 4, 2},
	})

	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}
}

func TestMinor(t *testing.T) {
	m := Matrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	output := m.Minor(1, 0)
	expected := 25.

	if !eq(output, expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}
}

func TestDeterminant3x3(t *testing.T) {
	m := Matrix([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	type test struct {
		row      int
		col      int
		expected float64
	}

	tds := []test{
		{0, 0, 56},
		{0, 1, 12},
		{0, 2, -46},
	}

	for i, td := range tds {
		output := m.Cofactor(td.row, td.col)
		if !eq(output, td.expected, epsilon) {
			t.Errorf("test %d failed: expected %f, returned %f", i, td.expected, output)
		}
	}

	d := m.Determinant()
	if !eq(d, -196, epsilon) {
		t.Errorf("expected -196 for determinant, returned %f", -d)
		return
	}
}

func TestDeterminant4x4(t *testing.T) {
	m := Matrix([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	type test struct {
		row      int
		col      int
		expected float64
	}

	tds := []test{
		{0, 0, 690},
		{0, 1, 447},
		{0, 2, 210},
		{0, 3, 51},
	}

	for i, td := range tds {
		output := m.Cofactor(td.row, td.col)
		if !eq(output, td.expected, epsilon) {
			t.Errorf("test %d failed: expected %f, returned %f", i, td.expected, output)
		}
	}

	d := m.Determinant()
	if !eq(d, -4071, epsilon) {
		t.Errorf("expected -4071 for determinant, returned %f", -d)
		return
	}
}

func TestInverse(t *testing.T) {
	m := Matrix([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})

	inverse, err := m.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	output := m.Multiply(inverse)
	expected := IdentityMatrix(4)
	if !output.Equal(expected, epsilon) {
		t.Errorf("expected %v, returned %v", expected, output)
		return
	}
}

func TestTranslation(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	m := TranslationMatrix(5, -3, 2)
	im, err := m.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	tts := []test{
		{m.MultiplyT(Point(-3, 4, 5)), Point(2, 1, 7)},
		{im.MultiplyT(Point(-3, 4, 5)), Point(-8, 7, 3)},
		{m.MultiplyT(Vector(-3, 4, 5)), Vector(-3, 4, 5)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}

func TestScaling(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	m := ScalingMatrix(2, 3, 4)
	im, err := m.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	tts := []test{
		{m.MultiplyT(Point(-4, 6, 8)), Point(-8, 18, 32)},
		{m.MultiplyT(Vector(-4, 6, 8)), Vector(-8, 18, 32)},
		{im.MultiplyT(Vector(-4, 6, 8)), Vector(-2, 2, 2)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}

func TestRotationX(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	fullQuarter := RotationXMatrix(math.Pi / 2.)
	halfQuarter := RotationXMatrix(math.Pi / 4.)
	ihalfQuarter, err := halfQuarter.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	tts := []test{
		{halfQuarter.MultiplyT(Point(0, 1, 0)), Point(0, math.Sqrt(2)/2., math.Sqrt(2)/2.)},
		{fullQuarter.MultiplyT(Point(0, 1, 0)), Point(0, 0, 1)},
		{ihalfQuarter.MultiplyT(Point(0, 1, 0)), Point(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}
func TestRotationY(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	fullQuarter := RotationYMatrix(math.Pi / 2.)
	halfQuarter := RotationYMatrix(math.Pi / 4.)
	ihalfQuarter, err := halfQuarter.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	tts := []test{
		{halfQuarter.MultiplyT(Point(0, 0, 1)), Point(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
		{fullQuarter.MultiplyT(Point(0, 0, 1)), Point(1, 0, 0)},
		{ihalfQuarter.MultiplyT(Point(0, 0, 1)), Point(-math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}

func TestRotationZ(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	fullQuarter := RotationZMatrix(math.Pi / 2.)
	halfQuarter := RotationZMatrix(math.Pi / 4.)
	ihalfQuarter, err := halfQuarter.Inverse(epsilon)
	if err != nil {
		t.Error(err)
		return
	}

	tts := []test{
		{halfQuarter.MultiplyT(Point(0, 1, 0)), Point(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
		{fullQuarter.MultiplyT(Point(0, 1, 0)), Point(-1, 0, 0)},
		{ihalfQuarter.MultiplyT(Point(0, 1, 0)), Point(math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}

func TestShearing(t *testing.T) {
	type test struct {
		input    Tuple
		expected Tuple
	}

	tts := []test{
		{ShearingMatrix(ShearingOptions{XpY: 1}).MultiplyT(Point(2, 3, 4)), Point(5, 3, 4)},
		{ShearingMatrix(ShearingOptions{XpZ: 1}).MultiplyT(Point(2, 3, 4)), Point(6, 3, 4)},
		{ShearingMatrix(ShearingOptions{YpX: 1}).MultiplyT(Point(2, 3, 4)), Point(2, 5, 4)},
		{ShearingMatrix(ShearingOptions{YpZ: 1}).MultiplyT(Point(2, 3, 4)), Point(2, 7, 4)},
		{ShearingMatrix(ShearingOptions{ZpX: 1}).MultiplyT(Point(2, 3, 4)), Point(2, 3, 6)},
		{ShearingMatrix(ShearingOptions{ZpY: 1}).MultiplyT(Point(2, 3, 4)), Point(2, 3, 7)},
	}

	for i, tt := range tts {
		if !tt.input.Equal(tt.expected, epsilon) {
			t.Errorf("test %d failed: expected %v, returned %v", i, tt.expected, tt.input)
		}
	}
}

func TestChained(t *testing.T) {
	A := RotationXMatrix(math.Pi / 2.)
	B := ScalingMatrix(5, 5, 5)
	C := TranslationMatrix(10, 5, 7)

	p1 := Point(1, 0, 1)
	p1 = A.MultiplyT(p1)
	p1 = B.MultiplyT(p1)
	p1 = C.MultiplyT(p1)

	p2 := Point(1, 0, 1)
	p2 = C.Multiply(B).Multiply(A).MultiplyT(p2)
	if !p1.Equal(p2, epsilon) {
		t.Errorf("expected %v, returned %v", p1, p2)
	}
}
