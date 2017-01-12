package slicer
import (
  "../geometry"
  "../utils"
  //"fmt"
  "strconv"
)

func Slice(filename string, model *geometry.Model, conf *utils.Config) {
  heighestz := model.HeighestZ()
  var linelist []geometry.LineSegment
  iterator := 0
  for sliceheight := float32(0.0); sliceheight <= heighestz; sliceheight += conf.LayerHeight {
    for key := range model.Triangles {
      if model.Triangles[key].IntersectsZ(sliceheight) {
        v1, v2 := model.Triangles[key].IntersectVectors(sliceheight)
        linelist = append(linelist, geometry.LineSegment{V1: v1, V2: v2})
        //fmt.Printf("Points Intersected: %v, %v\n", v1, v2)
      }
    }
    //fmt.Printf("%v", linelist)

    geometry.Save2DSlice(linelist, "layer_" + utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4) + ".png")
    linelist = nil
    iterator += 1
  }
}