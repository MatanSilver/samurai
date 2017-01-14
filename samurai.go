package main

import (
  "os"
  "github.com/urfave/cli"
  "./utils"
  "./parser"
  "./geometry"
  "./slicer"
  "fmt"
  "errors"
)

func main() {
  defaultconf := &utils.Config{LayerHeight: 0.2}
  app := cli.NewApp()
  app.Name = "samurai"
  app.Usage = "experimental stl slicer"
  app.Commands = []cli.Command{
    cli.Command{
      Name:   "slice",
      Usage:  "slice into gcode",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name:  "file, f",
          Usage: "Slice `FILE`",
        },
        cli.StringFlag{
          Name:  "config, c",
          Usage: "Load configuration from `FILE`",
        },
        cli.StringFlag{
          Name:   "output, o",
          Usage:  "Output gcode to `FILE`",
        },
      },
      Action: func(c *cli.Context) error {
        if (c.String("file") == "") {
          return errors.New("The -f/--file flag is required")
        }
        fmt.Printf("Loading triangles from STL file...\n")
        model := parser.ImportSTL(c.String("file"))
        fmt.Printf("%v triangles loaded\n", len(model.Triangles))
        var conf *utils.Config
        if c.String("config") != "" { //load config from flag, or default config
          conf = utils.LoadConfig(c.String("config"))
        } else {
          conf = defaultconf
        }
        output_name := "output.gcode"
        if (c.String("output") != "") { //load name from flag, or default name
          output_name = c.String("output")
        }
        slicer.Slice(output_name, model, conf)
        return nil
      },
    },
    cli.Command{
      Name:   "render",
      Usage:  "render stl to a png file",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name:     "file, f",
          Usage:    "Render `FILE`",
        },
        cli.StringFlag{
          Name:     "output, o",
          Usage:    "Output png to `FILE`",
        },
      },
      Action: func(c *cli.Context) error {
        fmt.Printf("Loading triangles from STL file...\n")
        model := parser.ImportSTL(c.String("file"))
        fmt.Printf("%v triangles loaded\n", len(model.Triangles))
        fmt.Printf("Rendering image...\n")
        geometry.Render(model, c.String("output"))
        fmt.Printf("Image rendered\n")
        return nil
      },
    },
    cli.Command{
      Name:   "generate_config",
      Usage:  "create a new config file from flags",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name:   "output, o",
          Usage:  "Output yaml config to `FILE`",
        },
        cli.StringFlag{
          Name:   "layer_height",
          Usage:  "Set the slice layer height",
        },
      },
      Action: func (c *cli.Context) error { //add ability to override defaults with flags
        output_name := "config.yaml"
        if (c.String("output") != "") {
          output_name = c.String("output")
        }
        fmt.Printf("Generating config file...\n")
        utils.GenerateConfig(output_name, &utils.Config{LayerHeight: 0.2})
        fmt.Printf("Config file generated\n")
        return nil
      },
    },
  }
  app.Run(os.Args)
}

