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

func (mat *Mat) Transpose() *Mat {
	matt := New(mat.Dimensions())
	for k := range mat.m {
		for j := range mat.m[k] {
			matt.m[j][k] = matt.m[k][j]
		}
	}
}
