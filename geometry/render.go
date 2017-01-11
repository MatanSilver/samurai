package geometry
import (
  "github.com/fogleman/ln/ln"
  //"fmt"
)


func Render(model *Model, filename string) {
// create a scene and add a single cube
    scene := ln.Scene{}
    for key := range model.Triangles {
      model.Triangles[key].Print()
      v1 := ln.Vector{float64(model.Triangles[key].Vertex1[0]), float64(model.Triangles[key].Vertex1[1]), float64(model.Triangles[key].Vertex1[2])}
      v2 := ln.Vector{float64(model.Triangles[key].Vertex2[0]), float64(model.Triangles[key].Vertex2[1]), float64(model.Triangles[key].Vertex2[2])}
      v3 := ln.Vector{float64(model.Triangles[key].Vertex3[0]), float64(model.Triangles[key].Vertex3[1]), float64(model.Triangles[key].Vertex3[2])}
      scene.Add(ln.NewTriangle(v1, v2, v3))
    }

    // define camera parameters
    eye := ln.Vector{160, 150, 20}    // camera position
    center := ln.Vector{100, 100, 0} // camera looks at
    up := ln.Vector{0, 0, 100}     // up direction

    // define rendering parameters
    width := 1024.0  // rendered width
    height := 1024.0 // rendered height
    fovy := 50.0     // vertical field of view, degrees
    znear := 0.1     // near z plane
    zfar := 200.0     // far z plane
    step := 0.01     // how finely to chop the paths for visibility testing

    // compute 2D paths that depict the 3D scene
    paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

    // render the paths in an image
    paths.WriteToPNG(filename, width, height)

    // save the paths as an svg
    //paths.WriteToSVG("out.svg", width, height)
}