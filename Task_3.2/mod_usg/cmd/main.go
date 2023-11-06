package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	yaml "github.com/go-yaml/yaml"

	cmd "results_usg/internal/cmd"
)

func main() {
	configFile, err := os.ReadFile("src/config.yaml")
	if err != nil {
		log.Fatalf("File reading error: %v", err)
	}

	//List of files' configs
	var config cmd.Config

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	if len(os.Args) > 1 {
		fmt.Println("Please, enter to search substring: ")
		reader := bufio.NewReader(os.Stdin)
		toSearchSubStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("String input error: %v", err)
		}
		toSearchSubStr = strings.TrimSpace(toSearchSubStr)

		isFound := false
		for _, fileConfig := range config.Files {
			if fileConfig.Filename == os.Args[1] && fileConfig.Substring == toSearchSubStr {
				isFound = true
				break
			}
		}
		_, err = os.Open(os.Args[1])
		if !isFound && err == nil {
			newFileConfig := cmd.FileConfig{
				Filename:  os.Args[1],
				Substring: toSearchSubStr,
			}

			config.Files = append(config.Files, newFileConfig)
			//configData - byte slice
			configData, err := yaml.Marshal(&config)
			if err != nil {
				log.Fatalf("Error marshaling config: %v\n", err)
			}
			//0644 - access rights
			if err := os.WriteFile("src/config.yaml", configData, 0644); err != nil {
				log.Fatalf("Error writing config file: %v\n", err)
			}
		}
	}

	a := cmd.NewApp()
	for _, fileConfig := range config.Files {
		isFind, err := a.Search(fileConfig.Filename, fileConfig.Substring)
		if err != nil {
			fmt.Printf("Processing file error %s: %v\n", fileConfig.Filename, err)
		} else if isFind {
			fmt.Printf("File %s contains substring %s\n", fileConfig.Filename, fileConfig.Substring)
		} else {
			fmt.Printf("File %s doesn't contain substring %s\n", fileConfig.Filename, fileConfig.Substring)
		}
	}
}
