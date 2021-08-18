package tracer

import (
	"errors"
	"math"
)

// Matrix represents a rectangular grid of numbers.
type Matrix [][]float64

// NewMatrix creates a new matrix of a specified length and height.
func NewMatrix(width, height int) Matrix {
	out := make([][]float64, height)
	for y := range out {
		out[y] = make([]float64, width)
	}
	return out
}

// IdentityMatrix returns the identity matrix of a specified size.
func IdentityMatrix(size int) Matrix {
	out := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		out[i][i] = 1
	}

	return out
}

// Equal returns true if a matrix is equal to the specified matrix.
func (m Matrix) Equal(m1 Matrix, e float64) bool {
	if m.width() != m1.width() || m.height() != m1.height() {
		return false
	}

	for i, row := range m {
		for j := range row {
			if !eq(m[i][j], m1[i][j], e) {
				return false
			}
		}
	}

	return true
}

// Multiply multiplies two matrices. Matrices must have equal width
// and height to be multiplied.
func (m Matrix) Multiply(m1 Matrix) Matrix {
	out := NewMatrix(m1.width(), m.height())

	for r := 0; r < out.height(); r++ {
		for c := 0; c < out.width(); c++ {
			for i := 0; i < out.height(); i++ {
				out[r][c] += m[r][i] * m1[i][c]
			}
		}
	}

	return out
}

// MultiplyT multiplies a matrix and a tuple. The matrix width must
// equal the length of the tuple to be multiplied.
func (m Matrix) MultiplyT(t Tuple) Tuple {
	out := Tuple(make([]float64, len(t)))

	for r := 0; r < m.height(); r++ {
		for i := 0; i < m.height(); i++ {
			out[r] += m[r][i] * t[i]
		}
	}

	return out
}

// Transpose transposes a matrix. Only square matrices can be transposed.
func (m Matrix) Transpose() Matrix {
	out := NewMatrix(m.height(), m.width())

	for r := 0; r < out.height(); r++ {
		for c := 0; c < out.width(); c++ {
			out[r][c] = m[c][r]
		}
	}

	return out
}

// Determinant calculates the determinate of a 2x2 matrix.
func (m Matrix) Determinant() float64 {
	if m.width() == 2 {
		return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	}

	out := 0.
	for i, val := range m[0] {
		out += val * m.Cofactor(0, i)
	}

	return out
}

// SubMatrix creates a smaller matrix from an existing matrix by removing
// a row and column.
func (m Matrix) SubMatrix(row, col int) Matrix {
	out := NewMatrix(m.width()-1, m.height()-1)

	rowIdx := 0
	for r := 0; r < m.height(); r++ {
		if r != row {

			colIdx := 0
			for c := 0; c < m.width(); c++ {
				if c != col {
					out[rowIdx][colIdx] = m[r][c]
					colIdx++
				}
			}
			rowIdx++
		}
	}

	return out
}

// Minor calculates the determinant of the submatrix specified.
func (m Matrix) Minor(row, col int) float64 {
	return m.SubMatrix(row, col).Determinant()
}

// Cofactor calculates the determinant of the submatrix specified, negating
// the value depending on which row and column is removed.
func (m Matrix) Cofactor(row, col int) float64 {
	out := m.Minor(row, col)
	if (row+col)%2 == 1 {
		out *= -1
	}
	return out
}

// Inverse inverts a matrix, failing if the matrix cannot be inverted.
func (m Matrix) Inverse(e float64) (Matrix, error) {
	d := m.Determinant()
	if eq(d, 0, e) {
		return nil, errors.New("matrix has determinant zero: cannot be inverted")
	}

	// create matrix of cofactor values divided by the determinant
	out := NewMatrix(m.width(), m.height())

	for r := 0; r < out.height(); r++ {
		for c := 0; c < out.width(); c++ {
			out[r][c] = m.Cofactor(r, c) / d
		}
	}

	return out.Transpose(), nil
}

// width is the width of a canvas.
func (m Matrix) width() int {
	return len(m[0])
}

// height is the height of a canvas.
func (m Matrix) height() int {
	return len(m)
}

// Note: the following functions apply to 4x4 matrices only.

// ScalingMatrix returns a 4x4 scaling matrix with the specified values.
func ScalingMatrix(x, y, z float64) Matrix {
	out := NewMatrix(4, 4)
	out[0][0] = x
	out[1][1] = y
	out[2][2] = z
	out[3][3] = 1

	return out
}

// TranslationMatrix returns a 4x4 translation matrix with the specified elements.
func TranslationMatrix(x, y, z float64) Matrix {
	out := IdentityMatrix(4)
	out[0][3] = x
	out[1][3] = y
	out[2][3] = z

	return out
}

// RotationX returns a 4x4 rotating matrix around the x-axis.
func RotationXMatrix(rad float64) Matrix {
	out := IdentityMatrix(4)

	out[1][1] = math.Cos(rad)
	out[1][2] = -1. * math.Sin(rad)
	out[2][1] = math.Sin(rad)
	out[2][2] = math.Cos(rad)

	return out
}

// RotationY returns a 4x4 rotating matrix around the y-axis.
func RotationYMatrix(rad float64) Matrix {
	out := IdentityMatrix(4)

	out[0][0] = math.Cos(rad)
	out[0][2] = math.Sin(rad)
	out[2][0] = -1. * math.Sin(rad)
	out[2][2] = math.Cos(rad)

	return out
}

// RotationZ returns a 4x4 rotating matrix around the z-axis.
func RotationZMatrix(rad float64) Matrix {
	out := IdentityMatrix(4)

	out[0][0] = math.Cos(rad)
	out[0][1] = -1. * math.Sin(rad)
	out[1][0] = math.Sin(rad)
	out[1][1] = math.Cos(rad)

	return out
}

// ShearingOptions represents the changing in proportion parameters
// of a shearing matrix (reads X in proportion to Y).
type ShearingOptions struct {
	XpY, XpZ, YpX, YpZ, ZpX, ZpY float64
}

// ShearingMatrix returns a 4x4 shearing matrix with the specified options.
func ShearingMatrix(opt ShearingOptions) Matrix {
	out := IdentityMatrix(4)

	out[0][1] = opt.XpY
	out[0][2] = opt.XpZ
	out[1][0] = opt.YpX
	out[1][2] = opt.YpZ
	out[2][0] = opt.ZpX
	out[2][1] = opt.ZpY

	return out
}
