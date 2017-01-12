package geometry
import (
  "fmt"
)


type Model struct {
  Header      [80]byte
  Count       uint32
  Triangles   []Triangle
  Length      uint32
}

func (m *Model) Print() {
  fmt.Printf("Model: \n\tHeader: %v\n\tCount: %v\n\tLength: %v\n\tTriangles: \n", m.Header, m.Count, m.Length)
  for key := range m.Triangles {
    m.Triangles[key].Print()
  }
}

func (m *Model) HeighestZ() float32 {
  var highest float32 = 0.0
  for key := range m.Triangles { //iterate triangles
    for key2 := range m.Triangles[key].Vertices { //iterate vertices
      if (m.Triangles[key].Vertices[key2][2] > highest) {
        highest = m.Triangles[key].Vertices[key2][2]
      }
    }
  }
  return highest
}

func (m *Model) GetTarget() Vector {
  return Vector{m.Triangles[0].Vertices[0][0], m.Triangles[0].Vertices[0][1], m.Triangles[0].Vertices[0][2]}
}

type Triangle struct {
  Normal      [3]float32
  Vertices    [3]Vector
  Attribute   uint16
}

func (tri *Triangle) Print() {
  fmt.Printf("Triangle\n\tNormal: %v\n\tVertex 1: %v\n\tVertex 2: %v\n\tVertex 3: %v\n\tAttribute: %v\n", tri.Normal, tri.Vertices[0], tri.Vertices[1], tri.Vertices[2], tri.Attribute)
}

func (tri *Triangle) IntersectsZ(zlevel float32) bool {
  if (tri.Vertices[0][2] > zlevel && tri.Vertices[1][2] > zlevel && tri.Vertices[2][2] > zlevel) {
    return false
  } else if (tri.Vertices[0][2] < zlevel && tri.Vertices[1][2] < zlevel && tri.Vertices[2][2] < zlevel) {
    return false
  }
  return true
}

type Vector [3]float32

func (vec *Vector) Cross(vec2 Vector) Vector {
  return Vector{vec[1]*vec2[2] - vec[2]*vec2[1], -(vec[0]*vec2[2]-vec[2]*vec2[0]), vec[0]*vec2[1]-vec[1]*vec2[0]}
}

func (vec *Vector) Dot(vec2 Vector) float32 {
  return float32(vec[0]*vec2[0] + vec[1]*vec2[1] + vec[2]*vec2[2])
}

func average(xs[]float32)float32 {
	total:=float32(0.0)
	for _,v:=range xs {
		total += v
	}
	return total/float32(len(xs))
}