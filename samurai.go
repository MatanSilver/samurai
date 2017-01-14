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
  conf := utils.DefaultConfig
  app := cli.NewApp()
  app.Name = "samurai"
  app.Usage = "experimental stl slicer"
  app.EnableBashCompletion = true
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
        cli.BoolFlag{
          Name:   "save_layer_images",
          Usage:  "Save rendered images of each layer",
        },
        cli.StringFlag{
          Name:   "layer_height",
          Usage:  "Set the slice layer height",
        },
        cli.IntFlag{
          Name:   "extruder_temp",
          Usage:  "Set extruder temperature",
          Destination:  &conf.ExtruderTemp,
        },
        cli.IntFlag{
          Name:   "bed_temp",
          Usage:  "Set bed temperature",
          Destination:  &conf.BedTemp,
        },
      },
      Action: func(c *cli.Context) error {
        if (c.String("file") == "") {
          return errors.New("The -f/--file flag is required")
        }
        fmt.Println("Loading triangles from STL file...")
        model := parser.ImportSTL(c.String("file"))
        fmt.Printf("%v triangles loaded\n", len(model.Triangles))
        if c.String("config") != "" { //load config from flag, or default config
          fmt.Println("Loading config...")
          conf = utils.LoadConfig(c.String("config"))
          fmt.Println("Config loaded")
        }
        output_name := "output.gcode"
        if (c.String("output") != "") { //load name from flag, or default name
          output_name = c.String("output")
        }
        fmt.Println("Slicing model...")
        slicer.Slice(output_name, model, conf, c.Bool("save_layer_images"))
        fmt.Println("Model sliced")
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
        if (c.String("file") == "") {
          return errors.New("The -f/--file flag is required")
        }
        output_name := "output.png"
        if (c.String("output") != "") {
          output_name = c.String("output")
        }
        fmt.Printf("Loading triangles from STL file...\n")
        model := parser.ImportSTL(c.String("file"))
        fmt.Printf("%v triangles loaded\n", len(model.Triangles))
        fmt.Printf("Rendering image...\n")
        geometry.Render(model, output_name)
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
        cli.IntFlag{
          Name:   "extruder_temp",
          Usage:  "Set extruder temperature",
          Destination:  &conf.ExtruderTemp,
        },
        cli.IntFlag{
          Name:   "bed_temp",
          Usage:  "Set bed temperature",
          Destination:  &conf.BedTemp,
        },
      },
      Action: func (c *cli.Context) error { //add ability to override defaults with flags
        output_name := "config.yaml"
        if (c.String("output") != "") {
          output_name = c.String("output")
        }
        fmt.Printf("Generating config file...\n")
        utils.GenerateConfig(output_name, conf)
        fmt.Printf("Config file generated\n")
        return nil
      },
    },
  }
  app.Run(os.Args)
}

