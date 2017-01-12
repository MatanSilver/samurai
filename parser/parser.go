package parser

import (
  "io/ioutil"
  "encoding/binary"
  "bytes"
  "os"
  "bufio"
  "strings"
  "strconv"
  "../utils"
  "../geometry"
)

func ImportSTL(filename string) *(geometry.Model) {
  file, err := os.Open(filename)
  utils.Check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  firstln := scanner.Text()
  var m *geometry.Model
  if strings.Contains(firstln, "solid") {
    m = ParseASCIISTL(filename)
  } else {
    m = ParseBinarySTL(filename)
  }
  //m.Print()
  return m
}

func ParseASCIISTL(filename string) *geometry.Model {
  file, err := os.Open(filename)
  utils.Check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  var triangles []geometry.Triangle
  var models []*geometry.Model
  var vertices []([3]float32)
  for scanner.Scan() {
    text := scanner.Text()
    if (strings.Contains(text, "solid") == true && strings.Contains(text, "endsolid") == false) {
      //fmt.Printf("new solid\n")
      models = append(models, new(geometry.Model))
    } else if strings.Contains(text, "facet normal") {
      //fmt.Printf("new facet normal\n")
      NormalStrings := strings.Split(strings.SplitAfter(text, "facet normal ")[1], " ")
      var NormalFloats [3]float32
      for i := 0; i < 3; i++ {
        Float64, err := strconv.ParseFloat(NormalStrings[i], 64)
        NormalFloats[i] = float32(Float64)
        utils.Check(err)
      }
      newtri := geometry.Triangle{Normal: NormalFloats}
      triangles = append(triangles, newtri)
    } else if strings.Contains(text, "vertex") {
      //fmt.Printf("new vertex\n")
      VertexStrings := strings.Split(strings.SplitAfter(text, "vertex ")[1], " ")
      var VertexFloats [3]float32
      for i := 0; i < 3; i++ {
        Float64, err := strconv.ParseFloat(VertexStrings[i], 32)
        VertexFloats[i] = float32(Float64)
        utils.Check(err)
      }
      vertices = append(vertices, VertexFloats)
    } else if strings.Contains(text, "endfacet") {
      //fmt.Printf("end facet\n")
      triangles[len(triangles) - 1].Vertices[0] = vertices[0]
      triangles[len(triangles) - 1].Vertices[1] = vertices[1]
      triangles[len(triangles) - 1].Vertices[2] = vertices[2]
      vertices = nil
    } else if strings.Contains(text, "endsolid") {
      //fmt.Printf("end solid\n")
      models[len(models) - 1].Triangles = triangles
      models[len(models) - 1].Length = uint32(len(triangles))
      triangles = nil
    }
  }
  return models[0]
}

func ParseBinarySTL(filename string) *geometry.Model {

  // Reading entire STL file into memory
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  m := new(geometry.Model)

  // Parsing Header first 80 bytes.
  err = binary.Read(bytes.NewBuffer(data[0:80]), binary.LittleEndian, &m.Header)
  if err != nil {
    panic(err)
  }

  // Parsing triangle count uint32 at byte 80
  err = binary.Read(bytes.NewBuffer(data[80:84]), binary.LittleEndian, &m.Length)
  if err != nil {
    panic(err)
  }

  // Allocating enough memory for all the triangles in the slice
  m.Triangles = make([]geometry.Triangle, m.Length)

  // Parsing the Triangle slice on byte 84 onwards, 50 bytes per triangle
  err = binary.Read(bytes.NewBuffer(data[84:]), binary.LittleEndian, &m.Triangles)
  if err != nil {
    panic(err)
  }

  return m
}