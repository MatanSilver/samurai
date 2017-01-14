package slicer
import (
  "../geometry"
  "../utils"
  "strconv"
  //"fmt"
)

func Slice(filename string, model *geometry.Model, conf *utils.Config) {
  //model.Print()
  heighestz := model.HeighestZ()
  linelists := [][]geometry.LineSegment{}
  iterator := 0
  for sliceheight := float32(0.0); sliceheight <= heighestz; sliceheight += conf.LayerHeight {
    linelists = append(linelists, []geometry.LineSegment{})
    for key := range model.Triangles {
      if model.Triangles[key].IntersectsZ(sliceheight) {
        //fmt.Printf("intersects")
        v1, v2 := model.Triangles[key].IntersectVectors(sliceheight)
        linelists[iterator] = append(linelists[iterator], geometry.LineSegment{V1: v1, V2: v2})
        //fmt.Printf("%v\n", linelists[iterator])
      }
    }
    //fmt.Printf("%v\n", linelists[iterator])
    if (len(linelists[iterator]) > 0){ //this was a weird thing
      geometry.Save2DSlice(linelists[iterator], "layer_" + utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4) + ".png")
    }
    iterator += 1
  }
}