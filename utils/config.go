package utils
import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
  //"fmt"
)

type Config struct {
  LayerHeight float32
  HeatBed bool
  BedTemp int
  ExtruderTemp int
}

func LoadConfig(filename string) Config {
  //dat, err := ioutil.ReadFile(filename)
  dat, err := ioutil.ReadFile(filename)
  Check(err)
  var conf Config
  err = yaml.Unmarshal([]byte(dat), &conf)
  Check(err)
  return conf
}

func GenerateConfig(filename string, conf Config) bool {
  dat, err := yaml.Marshal(conf)
  Check(err)
  err = ioutil.WriteFile(filename, dat, 0644)
  Check(err)
  return true
}

var DefaultConfig Config = Config{
  LayerHeight: 0.2,
  HeatBed: true,
  BedTemp: 60,
  ExtruderTemp: 220,
}