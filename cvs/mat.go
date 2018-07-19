package cvs

// A Mat is a matrix.
type Mat struct {
	m [][]complex128
}

// New creates a new matrix.
func New(rows, cols int) *Mat {
	mat := &Mat{
		m: make([][]complex128, cols),
	}

	for i := range mat.m {
		mat.m[i] = make([]complex128, rows)
	}

	return mat
}

// NewVec creates a new vector.
func NewVec(cols int) *Mat {
	return New(1, cols)
}

func NewVecFromInts(ns []int) *Mat {
	vec := NewVec(len(ns))
	for i := range ns {
		vec.m[i] = []complex128{complex(float64(ns[i]), 0)}
	}

	return vec
}

// check should perform any internal checks to make sure the matrix is consistent.
func (mat Mat) check() {
	if len(mat.m) == 0 {
		panic("can't check uninitialised matrix")
	}

	if len(mat.m) == 1 {
		return
	}

	m := len(mat.m[0])
	for i := range mat.m {
		if len(mat.m[i]) != m {
			panic("inconsistent matrix size")
		}
	}
}

// Dimensions returns the number of rows and columns in the matrix.
func (mat Mat) Dimensions() (int, int) {
	return len(mat.m[0]), len(mat.m)
}

// NewFrom creates a matrix by copying in an existing matrix of complex numbers.
func NewFrom(m [][]complex128) *Mat {
	mat := &Mat{
		m: m,
	}

	mat.check()
	return mat
}

// NewFromInts creates a matrix from an existing matrix of integers.
func NewFromInts(m [][]int) *Mat {
	mat := &Mat{m: make([][]complex128, len(m))}
	for i := range mat.m {
		mat.m[i] = make([]complex128, len(m[i]))
		for j := range mat.m[i] {
			mat.m[i][j] = complex(float64(m[i][j]), 0.0)
		}
	}
	mat.check()
	return mat
}

// Transpose computes the transpose of the matrix, returning a new matrix.
func (mat *Mat) Transpose() *Mat {
	matt := New(mat.Dimensions())
	for k := range mat.m {
		for j := range mat.m[k] {
			matt.m[j][k] = mat.m[k][j]
		}
	}

	return matt
}

func (mat *Mat) At(row, col int) complex128 {
	return mat.m[col][row]
}

// Eq? returns true if the two
func (mat *Mat) Eq(omat *Mat) bool {
	rows, cols := mat.Dimensions()
	orows, ocols := mat.Dimensions()
	if rows != orows || cols != ocols {
		return false
	}

	for k := 0; k < cols; k++ {
		for j := 0; j < rows; j++ {
			if mat.At(j, k) != omat.At(j, k) {
				return false
			}
		}
	}

	return true
}
