package slicer
import (
  "../geometry"
  "../utils"
  "strconv"
  "fmt"
  "os"
)

func Slice(filename string, model *geometry.Model, conf *utils.Config, save_layer_images bool) {
  //model.Print()
  //generate gcode for heating/inital things (homing, etc.)
  heighestz := model.HeighestZ()
  linelists := [][]geometry.LineSegment{}
  iterator := 0
  f, err := os.Create(filename)
  utils.Check(err)
  defer f.Close()
  _, err = f.WriteString(fmt.Sprintf("M107\nM190 S%v ; set bed temperature\nM104 S%v ; set temperature\nG28 ; home all axes\nG1 Z5 F5000 ; lift nozzle\nM109 S%v ; wait for temperature to be reached\nG21 ; set units to millimeters\nG90 ; use absolute coordinates\nM82 ; use absolute distances for extrusion\n", conf.BedTemp, conf.ExtruderTemp, conf.BedTemp))
  f.Sync()
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
    if (save_layer_images == true && len(linelists[iterator]) > 0){ //this was a weird thing
      geometry.Save2DSlice(linelists[iterator], "layer_" + utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4) + ".png")
    }
    //generate gcode for the layer here (plane and z change)
    iterator += 1
  }
}