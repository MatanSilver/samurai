package render
import (
  "github.com/fogleman/ln/ln"
  "samurai/geometry"
  //"sync"
  //"fmt"
  "github.com/llgcode/draw2d/draw2dimg"
  "image"
  "image/color"
  //"gopkg.in/cheggaaa/pb.v1"

)

type  RenderPayload struct {
  Scene *ln.Scene
  Triangle *geometry.Triangle
}

func AddTriangleToScene(scene *ln.Scene, triangle geometry.Triangle) {
  v1 := ln.Vector{float64(triangle.Vertices[0][0]), float64(triangle.Vertices[0][1]), float64(triangle.Vertices[0][2])}
  v2 := ln.Vector{float64(triangle.Vertices[1][0]), float64(triangle.Vertices[1][1]), float64(triangle.Vertices[1][2])}
  v3 := ln.Vector{float64(triangle.Vertices[2][0]), float64(triangle.Vertices[2][1]), float64(triangle.Vertices[2][2])}
  scene.Add(ln.NewTriangle(v1, v2, v3))
}


func Render(model geometry.Model, filename string) {
// create a scene and add a single cube
    scene := ln.Scene{}
    //count := len(model.Triangles)
    //bar := pb.StartNew(count)
    for key := range model.Triangles {
      //bar.Increment()
      AddTriangleToScene(&scene, model.Triangles[key])
    }
    //bar.FinishPrint("Complete")


    // define camera parameters
    eye := ln.Vector{200, 200, 200}    // camera position
    //center := ln.Vector{-300, -300, 0} // camera looks at
    target := model.GetTarget()
    center:= ln.Vector{float64(target[0]), float64(target[1]), float64(target[2])}
    up := ln.Vector{0, 0, 100}     // up direction

    // define rendering parameters
    width := 256.0  // rendered width
    height := 256.0 // rendered height
    fovy := 50.0     // vertical field of view, degrees
    znear := 0.1     // near z plane
    zfar := 700.0     // far z plane
    step := 0.01     // how finely to chop the paths for visibility testing

    // compute 2D paths that depict the 3D scene
    paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

    // render the paths in an image
    paths.WriteToPNG(filename, width, height)

    // save the paths as an svg
    //paths.WriteToSVG("out.svg", width, height)
}

func Save2DSlice(linelist []geometry.LineSegment, filename string) {
  dest := image.NewRGBA(image.Rect(0, 0, 200, 200))
  gc := draw2dimg.NewGraphicContext(dest)

  // Set some properties
  gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
  gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
  gc.SetLineWidth(1)

  xoffset := 50.0
  yoffset := 50.0
  xscale := 5.0
  yscale := 5.0
  //fmt.Printf("%v\n", linelist)
  gc.MoveTo(float64(linelist[0].V1[0])*xscale + xoffset, float64(linelist[0].V1[1])*yscale + yoffset) // should always be called first for a new path
  gc.LineTo(float64(linelist[0].V2[0])*xscale + xoffset, float64(linelist[0].V2[1])*yscale + yoffset)
  gc.Close()
  for key := range linelist[1:] {
    gc.MoveTo(float64(linelist[key].V1[0])*xscale + xoffset, float64(linelist[key].V1[1])*yscale + yoffset)
    gc.LineTo(float64(linelist[key].V2[0])*xscale + xoffset, float64(linelist[key].V2[1])*yscale + yoffset)
    gc.Close()
  }


  gc.FillStroke()
  // Save to file
  draw2dimg.SaveToPngFile(filename, dest)
}