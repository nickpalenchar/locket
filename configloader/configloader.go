package configloader

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Configopts struct {
	Version     string   `yaml:"files"`
	Directories []string `yaml:",flow"`
	Auth        struct {
		Aws struct {
			Profile string `yaml:"profile"`
			Bucket  string `yaml:"bucket"`
		}
	}
}

/* Config loads the config file located at the user's home */
func Config() *Configopts {
	opts := Configopts{}

	p := path.Join(os.Getenv("HOME"), ".locket.conf.yaml")

	var data, _ = os.ReadFile(p)

	err := yaml.Unmarshal([]byte(data), &opts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &opts
}
