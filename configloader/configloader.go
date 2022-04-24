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

/* Config loads the config file located at the user's home */
func Config() *Configopts {
	opts := Configopts{}

	p := path.Join(os.Getenv("HOME"), ".locket.conf.yaml")

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

func NewConfig(dir string, awsProfile string, awsBucket string) *Configopts {
	return &Configopts{
		Version: "1",
		Auth: struct{ Aws AwsOpts }{
			Aws: AwsOpts{
				Profile: awsProfile,
				Bucket:  awsBucket,
			},
		},
	}
}
