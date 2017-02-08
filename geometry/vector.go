package geometry

type Vector [3]float64

type Vector32 [3]float32

type LineSegment struct {
  V1 Vector
  V2 Vector
}

func (vec *Vector) Cross(vec2 Vector) Vector {
  return Vector{vec[1]*vec2[2] - vec[2]*vec2[1], -(vec[0]*vec2[2]-vec[2]*vec2[0]), vec[0]*vec2[1]-vec[1]*vec2[0]}
}

func (vec *Vector) Dot(vec2 Vector) float64 {
  return vec[0]*vec2[0] + vec[1]*vec2[1] + vec[2]*vec2[2]
}

func (vec *Vector) Rotate(rotation Vector, origin Vector) Vector {
return *vec
}

func (vec *Vector) RotateX(angle float64) Vector{
  //mat := Matrix{[]float64{1.0, 0.0, 0.0}, []float64{0, math.Cos(angle), -math.Sin(angle)}, []float64{0, math.Sin(angle), math.Cos(angle)}}
  return *vec
}