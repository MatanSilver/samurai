package slicer

import (
	"bufio"
	"fmt"
	"github.com/matansilver/samurai/geometry"
	"github.com/matansilver/samurai/render"
	"github.com/matansilver/samurai/utils"
	"os"
	"strconv"
	//"gopkg.in/cheggaaa/pb.v1"
)

func Slice(filename string, model *geometry.Model, conf utils.Config, save_layer_images bool) {
	//model.Print()
	//generate gcode for heating/inital things (homing, etc.)
	highestz := model.HighestZ()
	//linelists := [][]geometry.LineSegment{}
	linelists := []geometry.LineList{}
	iterator := 0
	corner := model.GetCornerVector()
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
	_, err = w.WriteString(fmt.Sprintf("G92 E0 ; reset extruder position\n"))
	utils.Check(err)
	w.Flush()
	//f.Sync()
	//count := int(highestz / conf.LayerHeight)
	//bar := pb.StartNew(count)
	for sliceheight := 0.0; sliceheight <= highestz; sliceheight += conf.LayerHeight {
		//bar.Increment()
		//linelists = append(linelists, []geometry.LineSegment{})
		linelists = append(linelists, geometry.LineList{})
		for key := range model.Triangles { //add lines from intersection of z with triangles
			if model.Triangles[key].IntersectsZ(sliceheight) {
				v1, v2 := model.Triangles[key].IntersectVectors(sliceheight)
				linelists[iterator] = append(linelists[iterator], geometry.LineSegment{V1: v1, V2: v2})
			}
		}
		if len(linelists[iterator]) != 0 {
		}
		if save_layer_images == true && len(linelists[iterator]) > 0 { //this was a weird thing
			os.Mkdir("layer_images", os.FileMode(0777))
			render.Save2DSlice(corner, linelists[iterator], "layer_images/layer_"+utils.LeftPad2Len(strconv.Itoa(iterator), "0", 4)+".png")
		}
		//make a list of linelists, each a complete closed loop
		openloops := geometry.LineListToOpenLoops(linelists[iterator])
		closedloops := geometry.CloseLoops(openloops)
		//panic("debug")
		//generate gcode for the layer here (plane and z change)
		epos := 0.0
		if sliceheight != 0.0 {
			_, err = w.WriteString(fmt.Sprintf("G1 Z%v\n", sliceheight))
			utils.Check(err)
			_, err = w.WriteString(fmt.Sprintf("G92 E0\n"))
			utils.Check(err)
		}
		for _, closedloop := range closedloops {
			utils.Check(err)
			for _, line := range closedloop {
				epos += line.Length() * 0.1
				_, err = w.WriteString(fmt.Sprintf("G1 X%v Y%v\n", line.V1[0], line.V1[1]))
				utils.Check(err)
				_, err = w.WriteString(fmt.Sprintf("G1 X%v Y%v E%v\n", line.V2[0], line.V2[1], epos))
				utils.Check(err)
			}
		}
		//_, err = w.WriteString(fmt.Sprintf("\n"))
		//utils.Check(err)
		//make shells

		//make interface

		//make infil

		//make support

		iterator += 1
		w.Flush()
	}
	//bar.FinishPrint("")
}
