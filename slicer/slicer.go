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

func Slice(filename string, model geometry.Model, conf utils.Config, save_layer_images bool) {
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
	w.Flush()
	//f.Sync()
	//count := int(highestz / conf.LayerHeight)
	//bar := pb.StartNew(count)
	for sliceheight := 5.0; sliceheight <= highestz; sliceheight += conf.LayerHeight {
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
		//seg1 := geometry.LineSegment{V1: geometry.Vector{0.0, 0.0, 0.0}, V2: geometry.Vector{0.0, 1.0, 0.0}}
		// seg1_alt := geometry.LineSegment{V2: geometry.Vector{0.0, 1.0, 0.0}, V1: geometry.Vector{0.0, 0.0, 0.0}}
		// seg2 := geometry.LineSegment{V1: geometry.Vector{0.0, 1.0, 0.0}, V2: geometry.Vector{1.0, 1.0, 0.0}}
		// seg3 := geometry.LineSegment{V1: geometry.Vector{1.0, 1.0, 0.0}, V2: geometry.Vector{1.0, 0.0, 0.0}}
		// seg4 := geometry.LineSegment{V1: geometry.Vector{1.0, 0.0, 0.0}, V2: geometry.Vector{0.0, 0.0, 0.0}}
		//ll := geometry.LineList{seg3, seg1_alt, seg2, seg4}
		// ll := geometry.LineList{seg4, seg1_alt, seg2, seg3}
		// openloops := geometry.LineListToOpenLoops(ll)
		closedloops := geometry.CloseLoops(openloops)
		fmt.Printf("closed loops: %v\n", closedloops)
		panic("debug")
		//fmt.Printf("%v\n", looplist)
		//generate gcode for the layer here (plane and z change)

		//_, err = w.WriteString(fmt.Sprintf("\n"))
		//utils.Check(err)
		//make shells

		//make interface

		//make infil

		//make support

		iterator += 1
	}
	//bar.FinishPrint("")
}
