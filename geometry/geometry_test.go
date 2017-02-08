package geometry

import (
  "testing"
  "fmt"
  "math"
)

func TestVector(t *testing.T) {

  if (!VectorEquals(Vector{1.0, 0.0, 0.0}, Vector{1.0, 0.0, 0.0})) {
    msg := fmt.Sprintf("Vectors %v, %v not equal", Vector{1.0, 0.0, 0.0}, Vector{1.0, 0.0, 0.0})
    t.Error(msg)
  }
  if (VectorEquals(Vector{1.0, 0.0, 0.0}, Vector{0.0, 0.0, 0.0})) {
    msg := fmt.Sprintf("Vectors %v, %v equal", Vector{1.0, 0.0, 0.0}, Vector{0.0, 0.0, 0.0})
    t.Error(msg)
  }
  if (!VectorEquals(*(&Vector{1.0, 0.0, 0.0}).Add(Vector{0.0, 1.0, 0.0}), Vector{1.0, 1.0, 0.0})) {
    t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 0.0, 0.0}).Add(Vector{0.0, 1.0, 0.0}), Vector{1.0, 1.0, 0.0}))
  }
  if (!VectorEquals(*(&Vector{1.0, 1.0, 0.0}).Subtract(Vector{0.0, 1.0, 0.0}), Vector{1.0, 0.0, 0.0})) {
    t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 1.0, 0.0}).Subtract(Vector{0.0, 1.0, 0.0}), Vector{1.0, 0.0, 0.0}))
  }
  if (!VectorEquals(*(&Vector{1.0, 0.0, 0.0}).Rotate(Vector{0.0, math.Pi/2, 0.0}, Vector{0.0, 0.0, 0.0}), Vector{0.0, 0.0, 1.0})) {
    t.Error(fmt.Sprintf("Vectors %v, %v not equal", *(&Vector{1.0, 0.0, 0.0}).Rotate(Vector{0.0, math.Pi/2, 0.0}, Vector{0.0, 0.0, 0.0}), Vector{0.0, 0.0, 1.0}))
  }
}