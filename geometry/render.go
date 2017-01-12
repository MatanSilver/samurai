package geometry
import (
  "github.com/fogleman/ln/ln"
  //"fmt"
)

func AddTriangleToScene(scene *ln.Scene, triangle *Triangle) {
  v1 := ln.Vector{float64(triangle.Vertices[0][0]), float64(triangle.Vertices[0][1]), float64(triangle.Vertices[0][2])}
  v2 := ln.Vector{float64(triangle.Vertices[1][0]), float64(triangle.Vertices[1][1]), float64(triangle.Vertices[1][2])}
  v3 := ln.Vector{float64(triangle.Vertices[2][0]), float64(triangle.Vertices[2][1]), float64(triangle.Vertices[2][2])}
  scene.Add(ln.NewTriangle(v1, v2, v3))
}


func Render(model *Model, filename string) {
// create a scene and add a single cube
    scene := ln.Scene{}
    for key := range model.Triangles {
      //model.Triangles[key].Print()
      AddTriangleToScene(&scene, &model.Triangles[key])
    }

    // define camera parameters
    eye := ln.Vector{0, 0, 300}    // camera position
    //center := ln.Vector{-300, -300, 0} // camera looks at
    target := model.GetTarget()
    center:= ln.Vector{float64(target[0]), float64(target[1]), float64(target[2])}
    up := ln.Vector{0, 0, 100}     // up direction

    // define rendering parameters
    width := 1024.0  // rendered width
    height := 1024.0 // rendered height
    fovy := 30.0     // vertical field of view, degrees
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