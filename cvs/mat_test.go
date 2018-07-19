package cvs

import "testing"

func TestNewMatrix(t *testing.T) {
	mat := New(10, 11)
	mat.check()
	m, n := mat.Dimensions()
	if m != 10 {
		t.Fatal("mat: expected 10 rows, have ", m)
	}
	if n != 11 {
		t.Fatal("mat: expected 11 columns, have ", n)
	}

	mat = NewFromInts([][]int{[]int{1,2,3}, []int{4,5,6}, []int{7,8,9}})
	m, n = mat.Dimensions()
	if m != 3 {
		t.Fatal("mat: expected 3 rows, have ", m)
	}
	if n != 3 {
		t.Fatal("mat: expected 3 columns, have ", n)
	}

	vec := NewVecFromInts([]int{1, 2, 3})
	rows, cols := vec.Dimensions()
	if rows != 1 && cols != 3 {
		t.Fatalf("mat: expected a 1×3 vector but have a %d×%d matrix", rows, cols)
	}
}
