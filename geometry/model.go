package geometry

import (
	"fmt"
)

type Model struct {
	Header    [80]byte
	Count     uint64
	Triangles []Triangle
	Length    uint64
}

type Model32 struct {
	Header    [80]byte
	Count     uint32
	Triangles []Triangle32
	Length    uint32
}

func (m *Model) Print() {
	fmt.Printf("Model: \n\tHeader: %v\n\tCount: %v\n\tLength: %v\n\tTriangles: \n", m.Header, m.Count, m.Length)
	for key := range m.Triangles {
		m.Triangles[key].Print()
	}
}

func (m *Model) HighestZ() float64 {
	var highest float64 = 0.0
	for key := range m.Triangles { //iterate triangles
		for key2 := range m.Triangles[key].Vertices { //iterate vertices
			if m.Triangles[key].Vertices[key2][2] > highest {
				highest = m.Triangles[key].Vertices[key2][2]
			}
		}
	}
	return highest
}

func (m *Model) GetTarget() Vector {
	return Vector{m.Triangles[0].Vertices[0][0], m.Triangles[0].Vertices[0][1], m.Triangles[0].Vertices[0][2]}
}

func (m *Model) GetCornerVector() Vector {
	leftmost := m.Triangles[0].Vertices[0]
	bottommost := m.Triangles[0].Vertices[0]
	for _, triangle := range m.Triangles {
		for _, vertex := range triangle.Vertices {
			if vertex[0] < leftmost[0] {
				leftmost = vertex
			}
			if vertex[1] < bottommost[1] {
				bottommost = vertex
			}
		}
	}
	return Vector{leftmost[0], bottommost[1], 0.0}
}

func (m *Model) Rotate(rot Vector, origin Vector) *Model {
	for key := range m.Triangles {
		m.Triangles[key].Rotate(rot, origin)
	}
	return m
}

func (m *Model) Translate(vec Vector) *Model {
	for key := range m.Triangles {
		m.Triangles[key].Translate(vec)
	}
	return m
}

func (m *Model) Jitter() *Model {
	origin := Vector{0.0, 0.0, 0.0}
	rot := Vector{0.00000001, 0.00000001, 0.00000001}
	m.Rotate(rot, origin)
	return m
}
