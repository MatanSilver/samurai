package geometry

import (
  "fmt"
)

type Triangle struct {
  Normal      Vector
  Vertices    [3]Vector
  Attribute   uint16
}

type Triangle32 struct {
  Normal      Vector32
  Vertices    [3]Vector32
  Attribute   uint16
}

func (tri *Triangle) Print() {
  fmt.Printf("Triangle\n\tNormal: %v\n\tVertex 1: %v\n\tVertex 2: %v\n\tVertex 3: %v\n\tAttribute: %v\n", tri.Normal, tri.Vertices[0], tri.Vertices[1], tri.Vertices[2], tri.Attribute)
}

func (tri *Triangle) IntersectsZ(zlevel float64) bool {
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

func (tri *Triangle) IntersectVectors(zlevel float64) (Vector, Vector) {
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

func (tri *Triangle) IntersectsLine(line LineSegment) bool {
  xterm := line.V2[0] - line.V1[0]
  yterm := line.V2[1] - line.V1[1]
  zterm := line.V2[2] - line.V1[2]
  p1 := tri.Vertices[0][0] * xterm + tri.Vertices[0][1] * yterm + tri.Vertices[0][2] * zterm
  p1sign := p1 >= 0
  p2 := tri.Vertices[1][0] * xterm + tri.Vertices[1][1] * yterm + tri.Vertices[1][2] * zterm
  p2sign := p2 >= 0
  p3 := tri.Vertices[2][0] * xterm + tri.Vertices[2][1] * yterm + tri.Vertices[2][2] * zterm
  p3sign := p3 >= 0
  if (p1sign && p2sign && p3sign || !p1sign && !p2sign && !p3sign) {
    return true
  }
  return false
}

func (tri *Triangle) Rotate(rot Vector, origin Vector) *Triangle {
  tri.Vertices[0].Rotate(rot, origin)
  tri.Vertices[1].Rotate(rot, origin)
  tri.Vertices[2].Rotate(rot, origin)
  return tri
}

func (tri *Triangle) Translate(vec Vector) *Triangle {
  tri.Vertices[0].Add(vec)
  tri.Vertices[1].Add(vec)
  tri.Vertices[2].Add(vec)
  return tri
}
