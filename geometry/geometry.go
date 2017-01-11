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

type Triangle struct {
  Normal      [3]float32
  Vertex1     [3]float32
  Vertex2     [3]float32
  Vertex3     [3]float32
  Attribute   uint16
}

func (tri *Triangle) Print() {
  fmt.Printf("Triangle\n\tNormal: %v\n\tVertex 1: %v\n\tVertex 2: %v\n\tVertex 3: %v\n\tAttribute: %v\n", tri.Normal, tri.Vertex1, tri.Vertex2, tri.Vertex3, tri.Attribute)
}