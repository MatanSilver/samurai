package main

import (
  "os"
  "bufio"
  "strings"
  "github.com/urfave/cli"
  "./utils"
  "./parser"
  "./geometry"
)

func main() {
  app := cli.NewApp()
  app.Name = "samurai"
  app.Usage = "experimental stl slicer"
  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name:  "file, f",
      Usage: "Slice `FILE`",
    },
    cli.StringFlag{
      Name:  "config, c",
      Usage: "Load configuration from `FILE`",
    },
    cli.StringFlag{
      Name:  "generate_config, g",
      Usage: "Generate configuration to `FILE`",
    },
    cli.StringFlag{
      Name:  "render, r",
      Usage: "Render image to `FILE`",
    },
    cli.StringFlag{
      Name:  "output, o",
      Usage: "Write gcode to `FILE`",
    },
  }

  app.Action = func(c *cli.Context) error {
    defaultconf := &utils.Config{LayerHeight: 0.2}
    if c.String("generate_config") != "" {
      utils.GenerateConfig(c.String("generate_config"), defaultconf)
    } else if c.String("file") != "" {
      model := ImportSTL(c.String("file"))
      if c.String("render") != "" {
        geometry.Render(model, c.String("render"))
      }
      if c.String("output") != "" {
        /*var conf *utils.Config
        if c.String("config") != "" {
          conf = utils.LoadConfig(c.String("config"))
        } else {
          conf = defaultconf
        }*/
        //slicer.Slice(c.String("output"), model, conf)
      }
    }
    return nil
  }
  app.Run(os.Args)
}

func ImportSTL(filename string) *(geometry.Model) {
  file, err := os.Open(filename)
  utils.Check(err)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  firstln := scanner.Text()
  var m *geometry.Model
  if strings.Contains(firstln, "solid") {
    m = parser.ParseASCIISTL(filename)
  } else {
    m = parser.ParseBinarySTL(filename)
  }
  //m.Print()
  return m
}

