package slicer
import (
  "../geometry"
  "../utils"
  "fmt"
)
func Slice(filename string, model *geometry.Model, conf *utils.Config) {
  heighestz := model.HeighestZ()
  for sliceheight := float32(0.0); sliceheight <= heighestz; sliceheight += conf.LayerHeight {
    fmt.Printf("Layer level: %v\n", sliceheight)
    var intersectingkeys []int
    for key := range model.Triangles {
      if (model.Triangles[key].IntersectsZ(sliceheight)) {
        intersectingkeys = append(intersectingkeys, key)
      }
    }
    fmt.Printf("Intersecting triangles: %v\n", intersectingkeys)
  }
}