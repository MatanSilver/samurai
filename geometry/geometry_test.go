package geometry

import (
	"fmt"
	"github.com/matansilver/samurai/utils"
	"math"
	"testing"
)

func TestVector(t *testing.T) {

	if (!VectorApprox(Vector{1.0, 0.0, 0.0}, Vector{1.0, 0.0, 0.0})) {
		msg := fmt.Sprintf("Vectors %v, %v not equal", Vector{1.0, 0.0, 0.0}, Vector{1.0, 0.0, 0.0})
		t.Error(msg)
	}
	if (VectorApprox(Vector{1.0, 0.0, 0.0}, Vector{0.0, 0.0, 0.0})) {
		msg := fmt.Sprintf("Vectors %v, %v equal", Vector{1.0, 0.0, 0.0}, Vector{0.0, 0.0, 0.0})
		t.Error(msg)
	}
	if (!VectorApprox(*(&Vector{1.0, 0.0, 0.0}).Add(Vector{0.0, 1.0, 0.0}), Vector{1.0, 1.0, 0.0})) {
		t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 0.0, 0.0}).Add(Vector{0.0, 1.0, 0.0}), Vector{1.0, 1.0, 0.0}))
	}
	if (!VectorApprox(*(&Vector{1.0, 1.0, 0.0}).Subtract(Vector{0.0, 1.0, 0.0}), Vector{1.0, 0.0, 0.0})) {
		t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 1.0, 0.0}).Subtract(Vector{0.0, 1.0, 0.0}), Vector{1.0, 0.0, 0.0}))
	}
	if (!VectorApprox(*(&Vector{1.0, 0.0, 0.0}).Rotate(Vector{0.0, math.Pi / 2, 0.0}, Vector{0.0, 0.0, 0.0}), Vector{0.0, 0.0, -1.0})) {
		t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 0.0, 0.0}).Rotate(Vector{0.0, math.Pi / 2, 0.0}, Vector{0.0, 0.0, 0.0}), Vector{0.0, 0.0, 1.0}))
	}
	if (!VectorApprox(VectorCross(Vector{1.0, 0.0, 0.0}, Vector{0.0, 1.0, 0.0}), Vector{0.0, 0.0, 1.0})) {
		msg := fmt.Sprintf("Vectors %v, %v not equal", VectorCross(Vector{1.0, 0.0, 0.0}, Vector{0.0, 1.0, 0.0}), Vector{0.0, 0.0, 1.0})
		t.Error(msg)
	}
	if math.Abs(VectorDot(Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0})-32.0) > 0.00001 {
		msg := fmt.Sprintf("%v . %v := %v", Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}, VectorDot(Vector{1.0, 2.0, 3.0}, Vector{4.0, 5.0, 6.0}))
		t.Error(msg)
	}
}

func TestGeometry(t *testing.T) {
	if average([]float64{1.0, 2.0, 3.0}) != 2.0 {
		msg := fmt.Sprintf("Averate of 1.0, 2.0, 3.0 was not 2.0")
		t.Error(msg)
	}
}

func TestMatrix(t *testing.T) {
	mat1 := Matrix{[]float64{1.0, 0.0, 0.0}, []float64{0.0, 1.0, 0.0}, []float64{0.0, 0.0, 1.0}}
	mat2 := Matrix{[]float64{1.0, 2.0, 3.0}}
	result := MatrixMultiply(mat1, mat2)
	truthy, err := MatrixApprox(result, mat2)
	utils.Check(err) //TODO is it idiomatic to check errors or throw through t.Error()?
	if truthy == false {
		msg := fmt.Sprintf("%vx%v is %v, not %v", mat1, mat2, mat2, result)
		t.Error(msg)
	}
}
