package geometry

import (
	"math"
	//"sync"
	"errors"
	//"fmt"
	"github.com/matansilver/samurai/utils"
)

type Vector [3]float64

type Vector32 [3]float32

type LineSegment struct {
	V1 Vector
	V2 Vector
}

type LineList []LineSegment

func (ls *LineSegment) Length() float64 {
	return ls.V1.SubtractVal(ls.V2).Magnitude()
}

func (ll *LineList) IsClosed() bool {
	if ll.IsOrdered() && VectorApprox((*ll)[0].V1, (*ll)[len(*ll)-1].V2) {
		return true
	}
	return false
}

func (ll *LineList) IsOrdered() bool {
	for key, line := range (*ll)[1:] { //TODO check if this works
		if !VectorApprox(line.V1, (*ll)[key].V2) {
			return false
		}
	}
	return true
}

func (ll *LineList) InsertLine(line LineSegment) bool {
	if VectorApprox(line.V2, (*ll)[0].V1) {
		*ll = append(LineList{line}, (*ll)...)
	} else if VectorApprox(line.V1, (*ll)[len(*ll)-1].V2) {
		*ll = append((*ll), line)
	} else {
		return false
	}
	return true
}

func (ll *LineList) InsertList(list LineList) (bool, error) {
	//newlist := LineList{}
	//fmt.Printf("trying to insert \n%v into \n%v\n", list, ll)
	if !list.IsOrdered() || !ll.IsOrdered() {
		return false, errors.New("One of the lists is not sorted")
	}
	if VectorApprox(list[len(list)-1].V2, (*ll)[0].V1) == true { //forwards in beginning
		//fmt.Printf("trying to insert forwards in beginning\n")
		*ll = append(list, (*ll)...)
	} else if VectorApprox(list[0].V1, (*ll)[len(*ll)-1].V2) == true { //forwards at end
		//fmt.Printf("trying to insert forwards in end\n")
		*ll = append((*ll), list...)
	} else if VectorApprox(list[0].V1, (*ll)[0].V1) == true { //backwards at beginning
		//fmt.Printf("trying to insert backwards in beginning\n")
		*ll = append(*list.FlipListVal(), (*ll)...)
	} else if VectorApprox(list[len(list)-1].V2, (*ll)[len(*ll)-1].V2) == true { //backwards at end
		//fmt.Printf("trying to insert backwards in end\n")
		*ll = append((*ll), *list.FlipListVal()...)
	} else {
		//fmt.Printf("not inserting\n")
		return false, nil
	}
	//fmt.Printf("resulted in: \n%v\n\n", ll)
	return true, nil
}

func VectorCross(vec1 Vector, vec2 Vector) Vector {
	return Vector{vec1[1]*vec2[2] - vec1[2]*vec2[1], -(vec1[0]*vec2[2] - vec1[2]*vec2[0]), vec1[0]*vec2[1] - vec1[1]*vec2[0]}
}

func VectorDot(vec1 Vector, vec2 Vector) float64 {
	return vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2]
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

func (vec *Vector) AddVal(vec2 Vector) *Vector {
	veccopy := *vec
	veccopy[0] += vec2[0]
	veccopy[1] += vec2[1]
	veccopy[2] += vec2[2]
	return &veccopy
}

func (vec *Vector) SubtractVal(vec2 Vector) *Vector {
	veccopy := *vec
	veccopy[0] -= vec2[0]
	veccopy[1] -= vec2[1]
	veccopy[2] -= vec2[2]
	return &veccopy
}

func (vec *Vector) Rotate(rot Vector, origin Vector) *Vector {
	return vec.Subtract(origin).RotateX(rot[0]).RotateY(rot[1]).RotateZ(rot[2]).Add(origin)
}

func (vec *Vector) RotateX(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{1.0, 0.0, 0.0}, []float64{0, math.Cos(angle), -math.Sin(angle)}, []float64{0, math.Sin(angle), math.Cos(angle)}}
	rotated_vec := MatrixMultiply([][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}}, rotation_mat)
	// vec[0] = rotated_vec[0][0]
	// vec[1] = rotated_vec[0][1]
	// vec[2] = rotated_vec[0][2]
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[1][0]
	vec[2] = rotated_vec[2][0]
	return vec
}

func (vec *Vector) RotateY(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{math.Cos(angle), 0.0, math.Sin(angle)}, []float64{0.0, 1.0, 0.0}, []float64{-math.Sin(angle), 0.0, math.Cos(angle)}}
	rotated_vec := MatrixMultiply([][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}}, rotation_mat)
	// vec[0] = rotated_vec[0][0]
	// vec[1] = rotated_vec[0][1]
	// vec[2] = rotated_vec[0][2]
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[1][0]
	vec[2] = rotated_vec[2][0]
	return vec
}

func (vec *Vector) RotateZ(angle float64) *Vector {
	rotation_mat := Matrix{[]float64{math.Cos(angle), -math.Sin(angle), 0.0}, []float64{math.Sin(angle), math.Cos(angle), 0.0}, []float64{0.0, 0.0, 1.0}}
	rotated_vec := MatrixMultiply([][]float64{[]float64{vec[0]}, []float64{vec[1]}, []float64{vec[2]}}, rotation_mat)
	// vec[0] = rotated_vec[0][0]
	// vec[1] = rotated_vec[0][1]
	// vec[2] = rotated_vec[0][2]
	vec[0] = rotated_vec[0][0]
	vec[1] = rotated_vec[1][0]
	vec[2] = rotated_vec[2][0]
	return vec
}

func VectorEquals(vec1 Vector, vec2 Vector) bool {
	return (vec1[0] == vec2[0] && vec1[1] == vec2[1] && vec1[2] == vec2[2])
}

func VectorApprox(vec1 Vector, vec2 Vector) bool {
	//fmt.Printf("comparing %v and %v\n", vec1, vec2)
	// mag := vec1.SubtractVal(vec2).Magnitude()
	// if mag == 0.0 {
	//  	fmt.Printf("returns: %v\n", (math.Abs(vec1[0]-vec2[0]) <= 0.00001 && math.Abs(vec1[1]-vec2[1]) <= 0.00001 && math.Abs(vec1[2]-vec2[2]) <= 0.00001))
	// }
	return (math.Abs(vec1[0]-vec2[0]) <= 0.00001 && math.Abs(vec1[1]-vec2[1]) <= 0.00001 && math.Abs(vec1[2]-vec2[2]) <= 0.00001)
}

func (vec *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec[0], 2) + math.Pow(vec[1], 2) + math.Pow(vec[2], 2))
}

func (ls *LineSegment) FlipSegment() *LineSegment {
	buff := ls.V1
	ls.V1 = ls.V2
	ls.V2 = buff
	return ls
}

func (ls *LineSegment) FlipSegmentVal() LineSegment {
	ls2 := *ls
	buff := ls.V1
	ls2.V1 = ls.V2
	ls2.V2 = buff
	return ls2
}

func (ll *LineList) FlipList() *LineList {
	//fmt.Printf("flipping!\n")
	for i := 0; i < len(*ll)/2; i++ {
		j := len(*ll) - i - 1
		(*ll)[i], (*ll)[j] = (*ll)[j], (*ll)[i]
	}
	for key := range *ll {
		(*ll)[key].FlipSegment()
	}
	return ll
}

func (ll *LineList) FlipListVal() *LineList {
	//fmt.Printf("flipping!\n")
	ll2 := *ll
	for i := 0; i < len(ll2)/2; i++ {
		j := len(ll2) - i - 1
		ll2[i], ll2[j] = ll2[j], ll2[i]
	}
	for key := range ll2 {
		ll2[key].FlipSegment()
	}
	return &ll2
}

func LineListToOpenLoops(linelist LineList) ([]LineList, error) {
	openloops := []LineList{}
	for _, val := range linelist {
		openloops = append(openloops, LineList{val})
	}
	return openloops, nil
}

func CloseLoops(openloops []LineList) ([]LineList, error) {
	closedloops := []LineList{}
	tries := 0
	for len(openloops) > 0 {
		iters := len(openloops) - 1
		//fmt.Printf("iters: %d\n", iters)
		//fmt.Printf("len(openloops): %d\n", len(openloops))
		for i := 0; i < iters; i++ { //iterate over all but last
			//fmt.Printf("i: %d\n", i)
			//fmt.Printf("trying to insert list of indices %d and %d\n", i, len(openloops)-1);
			//fmt.Printf("lists are: %v, and %v\n", openloops[i], openloops[len(openloops)-1])
			result, err := openloops[i].InsertList(openloops[len(openloops)-1])
			utils.Check(err)
			if result == true { //insert current and last element
				tries = 0
				openloops = openloops[0 : len(openloops)-1] //cut off last element
				//fmt.Printf("stitched together two lists\n")
				//fmt.Printf("openloops reduced to %d elements\n", len(openloops))
				iters--
				if openloops[i].IsClosed() {
					//fmt.Printf("Closed a loop\n")
					closedloops = append(closedloops, openloops[i])
					i--
					if len(openloops) > 1 {
						//fmt.Printf("i+1: %d\n", i+1)
						//fmt.Printf("len(openloops)-1: %d\n", len(openloops)-1)
						openloops = append(openloops[:i], openloops[i+1:len(openloops)-1]...) //TODO: might have to check for end of slice
					} else {
						openloops = []LineList{}
					}
					iters--
				}
			} else {
				//fmt.Printf("missed stitching\n")
				if tries > 50 {
					return closedloops, errors.New("unable to stitch loops")
				}
			}
			tries++
		}
	}
	return closedloops, nil
}
