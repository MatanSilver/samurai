package geometry

import (
  "errors"
)

type Matrix [][]float64

func MatrixMultiply(mat1 Matrix, mat2 Matrix) Matrix {
  var mat3 Matrix
  if (len(mat1) != len(mat2[0])) {
    err := errors.New("Matrix dimension mismatch")
    panic(err)
  } else {
    mat3 = make(Matrix, len(mat2))
    for i := 0; i < len(mat2); i++ {
      mat3[i] = make([]float64, len(mat1[0]))
      for j := 0; j < len(mat1[0]); j++ { //this iterates the mat3 location

        for k := 0; k < len(mat1); k++ { //iterates the multiplication location
            mat3[i][j] += mat1[k][j] * mat2[i][k]
        }

      }
    }
  }
  return mat3
}