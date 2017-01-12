package utils
import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Config struct {
  LayerHeight float32
}

func LoadConfig(filename string) *Config {
  //dat, err := ioutil.ReadFile(filename)
  dat, err := ioutil.ReadFile(filename)
  check(err)
  err := yaml.Marshal(conf)
  var conf Config
  err := yaml.Unmarshal([]byte(dat), &conf)
  return &conf
}

func GenerateConfig(filename string, conf *Config) bool {
  dat, err := yaml.Marshal(conf)
  Check(err)
  err = ioutil.WriteFile(filename, dat, 0644)
  Check(err)
  return true
}
