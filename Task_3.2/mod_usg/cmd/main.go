package main

import (
	"fmt"
	"log"
	"os"

	search "github.com/Bazhenator/StringFuncs/pkg/strSearch"
	yaml "github.com/go-yaml/yaml"

	cmd "results_usg/internal/cmd"
)

func main() {
	configFile, err := os.ReadFile("src/config.yaml")
	if err != nil {
		log.Fatalf("File reading error: %v", err)
	}

	var config cmd.Config

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	for _, fileConfig := range config.Files {
		isFind, err := search.Contains(fileConfig.Filename, fileConfig.Substring)
		if err != nil {
			log.Fatalf("Processing file error %s: %v\n", fileConfig.Filename, err)
		} else if isFind {
			fmt.Printf("File %s contains substring %s\n", fileConfig.Filename, fileConfig.Substring)
		} else {
			fmt.Printf("File %s doesn't contain substring %s\n", fileConfig.Filename, fileConfig.Substring)
		}
	}
}
