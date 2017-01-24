package slicer
import (
  "samurai/geometry"
  "samurai/utils"
  "strconv"
  "fmt"
  "os"
)

func Slice(filename string, model geometry.Model, conf utils.Config, save_layer_images bool) {
  //model.Print()
  //generate gcode for heating/inital things (homing, etc.)
  heighestz := model.HeighestZ()
  linelists := [][]geometry.LineSegment{}
  iterator := 0
  f, err := os.Create(filename)
  defer f.Close()
  utils.Check(err)
  w := bufio.NewWriter(f)
  _, err = w.WriteString(fmt.Sprintf("M107\n"))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("M190 S%v ; set bed temperature\n", conf.BedTemp))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("M104 S%v ; set temperature\n", conf.ExtruderTemp))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("G28 ; home all axes\n"))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("G1 Z5 F5000 ; lift nozzle\n"))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("M109 S%v ; wait for temperature to be reached\n", conf.BedTemp))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("G21 ; set units to millimeters\n"))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("G90 ; use absolute coordinates\n"))
  utils.Check(err)
  _, err = w.WriteString(fmt.Sprintf("M82 ; use absolute distances for extrusion\n"))
  utils.Check(err)
  w.Flush()
  //f.Sync()
  for sliceheight := float32(0.0); sliceheight <= heighestz; sliceheight += conf.LayerHeight {
    linelists = append(linelists, []geometry.LineSegment{})
    for key := range model.Triangles {
      if model.Triangles[key].IntersectsZ(sliceheight) {
        v1, v2 := model.Triangles[key].IntersectVectors(sliceheight)
        linelists[iterator] = append(linelists[iterator], geometry.LineSegment{V1: v1, V2: v2})
      }
    }
    if (len(linelists[iterator]) != 0) {
    }
    if (save_layer_images == true && len(linelists[iterator]) > 0){ //this was a weird thing
      geometry.Save2DSlice(linelists[iterator], "layer_" + utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4) + ".png")
    }
    //generate gcode for the layer here (plane and z change)
    iterator += 1
  }
}