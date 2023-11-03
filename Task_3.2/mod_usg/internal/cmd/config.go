package cmd

type FileConfig struct {
	Filename  string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

type Config struct {
	Files []FileConfig `yaml:"files"`
}
