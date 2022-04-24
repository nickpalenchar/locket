package metadata

type LocketMetadata struct {
	ApiVersion string
	ConfigFile string
}

func Metadata() *LocketMetadata {
	return &LocketMetadata{
		ApiVersion: "1",
		ConfigFile: ".locket.conf.yaml",
	}
}

func ApiVersion() string {
	return Metadata().ApiVersion
}

func ConfigFile() string {
	return Metadata().ConfigFile
}
