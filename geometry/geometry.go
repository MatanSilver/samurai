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
  z1 := tri.Vertices[0][2]
  z2 := tri.Vertices[1][2]
  z3 := tri.Vertices[2][2]
  if (z1 > zlevel && z2 > zlevel && z3 > zlevel) {
    return false
  } else if (z1 < zlevel && z2 < zlevel && z3 < zlevel) {
    return false
  } else if (z1 == zlevel && z2 == zlevel && z3 == zlevel) {
    return false
  } else if (z1 == zlevel && z2 > zlevel && z3 > zlevel) ||
  (z1 == zlevel && z2 < zlevel && z3 < zlevel) ||
  (z2 == zlevel && z1 > zlevel && z3 > zlevel) ||
  (z2 == zlevel && z1 < zlevel && z3 < zlevel) ||
  (z3 == zlevel && z1 > zlevel && z3 > zlevel) ||
  (z3 == zlevel && z1 < zlevel && z3 < zlevel) {
      return false
    }
  return true
}

func (tri *Triangle) IntersectVectors(zlevel float32) (Vector, Vector) {
  //if two points on z level, return those points
  if (tri.Vertices[0][2] == zlevel && tri.Vertices[1][2] == zlevel) {
    return tri.Vertices[0], tri.Vertices[1]
  } else if (tri.Vertices[1][2] == zlevel && tri.Vertices[2][2] == zlevel) {
    return tri.Vertices[1], tri.Vertices[2]
  } else if (tri.Vertices[0][2] == zlevel && tri.Vertices[2][2] == zlevel) {
    return tri.Vertices[0], tri.Vertices[2]
  } else { // otherwise, calculate the two intersections parametrically (could do rref?)
    var va, vb, origin Vector
    if (tri.Vertices[0][2] <= zlevel && tri.Vertices[1][2] <= zlevel && tri.Vertices[2][2] > zlevel) ||
    (tri.Vertices[0][2] >= zlevel && tri.Vertices[1][2] >= zlevel && tri.Vertices[2][2] < zlevel) {
      va = tri.Vertices[0]
      vb = tri.Vertices[1]
      origin = tri.Vertices[2]
    } else if (tri.Vertices[1][2] <= zlevel && tri.Vertices[2][2] <= zlevel && tri.Vertices[0][2] > zlevel) ||
    (tri.Vertices[1][2] >= zlevel && tri.Vertices[2][2] >= zlevel && tri.Vertices[0][2] < zlevel) {
      va = tri.Vertices[1]
      vb = tri.Vertices[2]
      origin = tri.Vertices[0]
    } else if (tri.Vertices[0][2] <= zlevel && tri.Vertices[2][2] <= zlevel && tri.Vertices[1][2] > zlevel) ||
    (tri.Vertices[0][2] >= zlevel && tri.Vertices[2][2] >= zlevel && tri.Vertices[1][2] < zlevel) {
      va = tri.Vertices[0]
      vb = tri.Vertices[2]
      origin = tri.Vertices[1]
    }
    t1 := (zlevel - va[2])/(origin[2] - va[2])
    t2 := (zlevel - vb[2])/(origin[2] - vb[2])
    x1 := va[0] + t1*(origin[0] - va[0])
    x2 := vb[0] + t2*(origin[0] - vb[0])
    y1 := va[1] + t1*(origin[1] - va[1])
    y2 := vb[1] + t2*(origin[1] - vb[1])
    v1 := Vector{x1, y1, zlevel}
    v2 := Vector{x2, y2, zlevel}
    return v1, v2
  }
}

type Vector [3]float32

type LineSegment struct {
  V1 Vector
  V2 Vector
}

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