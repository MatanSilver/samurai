package slicer
import (
  "../geometry"
  "../utils"
  "strconv"
)

func Slice(filename string, model *geometry.Model, conf *utils.Config) {
  heighestz := model.HeighestZ()
  linelists := [][]geometry.LineSegment{}
  iterator := 0
  for sliceheight := float32(0.0); sliceheight <= heighestz; sliceheight += conf.LayerHeight {
    linelists = append(linelists, []geometry.LineSegment{})
    for key := range model.Triangles {
      if model.Triangles[key].IntersectsZ(sliceheight) {
        v1, v2 := model.Triangles[key].IntersectVectors(sliceheight)
        linelists[iterator] = append(linelists[iterator], geometry.LineSegment{V1: v1, V2: v2})
      }
    }
    geometry.Save2DSlice(linelists[iterator], "layer_" + utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4) + ".png")
    iterator += 1
  }
}