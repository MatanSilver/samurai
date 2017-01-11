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
  return &Config{LayerHeight: 0.2} //TODO make this actually load from yaml
}

func GenerateConfig(filename string, conf *Config) bool {
  dat, err := yaml.Marshal(conf)
  Check(err)
  err = ioutil.WriteFile(filename, dat, 0644)
  Check(err)
  return false //TODO make this actually save a config yaml
}
