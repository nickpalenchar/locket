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
		Aws AwsOpts
	}
}

type AwsOpts struct {
	Profile string `yaml:"profile"`
	Bucket  string `yaml:"bucket"`
}

func (c *Configopts) ToFile(filepath string) {
	yamlData, err := yaml.Marshal(c)

	if err != nil {
		log.Fatalf("Error writing config to file: %s", err)
	}

	os.WriteFile(filepath, yamlData, 0666)
}

/* Config loads the config file located at the user's home */
func Config() *Configopts {
	opts := Configopts{}

	p := ConfigPath()

	if _, err := os.Stat(p); err != nil {
		log.Fatalf("error: No config file found (looking for %s)", p)
	}

	var data, _ = os.ReadFile(p)

	err := yaml.Unmarshal([]byte(data), &opts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &opts
}

func ConfigPath() string {
	return path.Join(os.Getenv("HOME"), ".locket.conf.yaml")
}

func NewConfig(dir string, awsProfile string, awsBucket string) *Configopts {
	return &Configopts{
		Version:     "1",
		Directories: []string{dir},
		Auth: struct{ Aws AwsOpts }{
			Aws: AwsOpts{
				Profile: awsProfile,
				Bucket:  awsBucket,
			},
		},
	}
}
