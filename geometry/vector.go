package geometry

import (
	"math"
	"sync"
	"errors"
)

type Vector [3]float64

type Vector32 [3]float32

type LineSegment struct {
	&sync.RWMutex{}
	V1 Vector
	V2 Vector
}

type LineList struct {
		&sync.RWMutex{}
		List []LineSegment
}

func (ll *LineList) IsClosed() bool {
	if (ll.IsOrdered() && VectorApprox(ll.list[0].V1, ll.list[-1].V2)) {
		return true
	}
	return false
}

func (ll *LineList) IsOrdered() bool {
	for key, line := range ll.List[1:] { //TODO check if this works
		if !VectorApprox(line.V1, ll.List[key].V2) {
			return false
		}
	}
	return true
}

func (ll *LineList) InsertLine(line LineSegment) bool {
	if VectorApprox(line.V2, ll.List[0].V1) {
		ll.List = append(line, ll.List...)
	} else if VectorApprox(line.V1, ll.List[-1].V2) {
		ll.list = append(ll.List..., line)
	} else {
		return false
	}
	return true
}

func (ll *LineList) InsertList(list LineList) (bool, error) {
	if VectorApprox(list.List[-1].V2, ll.List[0].V1) {
		if (!list.IsOrdered() || !ll.IsOrdered()) {
			return false, errors.New("One of the lists is not sorted")
		}
		ll.List = append(list..., ll.List...)
	} else if VectorApprox(list.List[0].V1, ll.List[-1].V2) {
		if (!list.IsOrdered() || !ll.IsOrdered()) {
			return false, errors.New("One of the lists is not sorted")
		}
		ll.list = append(ll.List..., list...)
	} else {
		return false, nil
	}
	return true, nil
}

func (vec *Vector) Cross(vec2 Vector) Vector {
	return Vector{vec[1]*vec2[2] - vec[2]*vec2[1], -(vec[0]*vec2[2] - vec[2]*vec2[0]), vec[0]*vec2[1] - vec[1]*vec2[0]}
}

func (vec *Vector) Dot(vec2 Vector) float64 {
	return vec[0]*vec2[0] + vec[1]*vec2[1] + vec[2]*vec2[2]
}

func (vec *Vector) Add(vec2 Vector) *Vector {
	vec[0] += vec2[0]
	vec[1] += vec2[1]
	vec[2] += vec2[2]
	return vec
}

func (vec *Vector) Subtract(vec2 Vector) *Vector {
	vec[0] -= vec2[0]
	vec[1] -= vec2[1]
	vec[2] -= vec2[2]
	return vec
}

func (vec *Vector) Rotate(rot Vector, origin Vector) *Vector {
	return vec.Subtract(origin).RotateX(rot[0]).RotateY(rot[1]).RotateZ(rot[2]).Add(origin)
}

func (vec *Vector) RotateX(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{1.0, 0.0, 0.0}, []float64{0, math.Cos(angle), -math.Sin(angle)}, []float64{0, math.Sin(angle), math.Cos(angle)}}
	rotated_vec := MatrixMultiply(rotation_mat, [][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}})
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[0][1]
	vec[2] = rotated_vec[0][2]
	return vec
}

func (vec *Vector) RotateY(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{math.Cos(angle), 0.0, math.Sin(angle)}, []float64{0.0, 1.0, 0.0}, []float64{-math.Sin(angle), 0.0, math.Cos(angle)}}
	rotated_vec := MatrixMultiply(rotation_mat, [][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}})
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[0][1]
	vec[2] = rotated_vec[0][2]
	return vec
}

func (vec *Vector) RotateZ(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{math.Cos(angle), -math.Sin(angle), 0.0}, []float64{math.Sin(angle), math.Cos(angle), 0.0}, []float64{0.0, 0.0, 1.0}}
	rotated_vec := MatrixMultiply(rotation_mat, [][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}})
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[0][1]
	vec[2] = rotated_vec[0][2]
	return vec
}

func VectorEquals(vec1 Vector, vec2 Vector) bool {
	return (vec1[0] == vec2[0] && vec1[1] == vec2[1] && vec1[2] == vec2[2])
}

func VectorApprox(vec1 Vector, vec2 Vector) bool {
	return (math.Abs(vec1[0]-vec2[0]) <= 0.000001 && math.Abs(vec1[1]-vec2[2]) <= 0.000001 && math.Abs(vec1[2]-vec2[2]) <= 0.00001)
}

func (ls *LineSegment) FlipSegment() *LineSegment {
	buff := ls.V1
	ls.V1 = ls.V2
	ls.V2 = buff
	return ls
}

func (ls *LineSegment) FlipSegmentVal() LineSegment {
	ls2 := ls
	buff := ls.V1
	ls2.V1 = ls.V2
	ls2.V2 = buff
	return ls2
}

/*
inputs:
line sebments
open loops

output:
open loops
*/
func segment_worker(chin <-chan geometry.LineSegment, chout chan<- []geometry.LineSegment) {

}

/*
input:
linelists
open loops

output:
open loops
closed loops
*/
func loop_worker(chin <-chan []geometry.LineSegment) {

}

func LineListToLoops(linelist LineList) [][]LineSegment {
	loops := [][]LineSegment{}
	loops = append(loops, []LineSegment{})
	semiloops := [][]LineSegment{}
	semiloops = append(loops, []LineSegment{})
	semiloopsmtx = &sync.RWMutex{}
	loops[0] = append(loops[0], linelist[0])
	semiloops[0] = append(semiloops[0], linelist[0])
	for key, line := range linelist[1:] {
		go func() {
			if VectorApprox(line.V1, linelist[key].V2) { //correct dir
				semiloops[len(semiloops)-1] = append(semiloops[len(semiloops)-1], line)
			} else if VectorApprox(line.V2, linelist[key].V2) { //wrong dir: flip current
				semiloops[len(semiloops)-1] = append(semiloops[len(semiloops)-1], line.FlipSegment())
			} else { //line too far: make new semiloops
				loops = append(semiloops, []geometry.LineSegment{})
				semiloops[len(semiloops)-1] = append(semiloops[len(semiloops)-1], line)
			}
		}()
	}
	return loops
}
