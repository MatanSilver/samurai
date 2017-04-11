package geometry

import (
	"errors"
	//"fmt"
	"math"
)

type Matrix [][]float64

func MatrixMultiply(mat1 Matrix, mat2 Matrix) Matrix {
	var mat3 Matrix
	if len(mat1) != len(mat2[0]) {
		err := errors.New("Matrix dimension mismatch")
		panic(err)
	} else {
		mat3 = make(Matrix, len(mat2))
		for i := 0; i < len(mat2); i++ {
			mat3[i] = make([]float64, len(mat1[0]))
			for j := 0; j < len(mat1[0]); j++ { //this iterates the mat3 location
				for k := 0; k < len(mat1); k++ { //iterates the multiplication location
					// fmt.Printf("mat1: %dx%d\n", len(mat1), len(mat1[0]))
					// fmt.Printf("mat2: %dx%d\n", len(mat2), len(mat2[0]))
					// fmt.Printf("mat3: %dx%d\n", len(mat3), len(mat3[0]))
					// fmt.Printf("mat3[%d][%d](%f) += mat1[%d][%d](%f) * mat2[%d][%d](%f)\n", i, j, mat3[i][j], k, j, mat1[k][j], i, k, mat2[i][k])
					//a, b := mat1[k][j], mat2[i][k]
					mat3[i][j] += mat1[k][j] * mat2[i][k]
					//mat3[i][j] += a * b
				}

			}
		}
	}
	return mat3
}

func MatrixEquals(mat1 Matrix, mat2 Matrix) (bool, error) {
	var err error
	//check dimensions
	if len(mat1) != len(mat2) || len(mat1[0]) != len(mat2[0]) {
		err = errors.New("Matrix dimension mismatch")
		return false, err
	} else {
		for i := range mat1 {
			for j := range mat1[0] {
				if mat1[i][j] != mat2[i][j] {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func MatrixApprox(mat1 Matrix, mat2 Matrix) (bool, error) {
	var err error
	//check dimensions
	if len(mat1) != len(mat2) || len(mat1[0]) != len(mat2[0]) {
		err = errors.New("Matrix dimension mismatch")
		return false, err
	} else {
		for i := range mat1 {
			for j := range mat1[0] {
				if math.Abs(mat1[i][j]-mat2[i][j]) > 0.00001 {
					return false, nil
				}
			}
		}
	}
	return true, nil
}
