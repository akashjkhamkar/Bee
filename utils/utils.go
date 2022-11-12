package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Yaml_config struct {
	Function_name string;
	Runtime string;
	Replicas int;
	Autoscaling bool;
}

func Read_yaml_config_file(filename string) Yaml_config {
	yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Error reading YAML file: %s\n", err)
    }

    var yamlConfig Yaml_config
    err = yaml.Unmarshal(yamlFile, &yamlConfig)

	if err != nil {
        fmt.Printf("Error parsing YAML file: %s\n", err)
    }

	return yamlConfig
}

func Create_yaml_config_file(data Yaml_config, filename string) {
	yamlData, err := yaml.Marshal(&data)

    if err != nil {
        fmt.Printf("Error while Marshaling. %v", err)
    }

    err = ioutil.WriteFile(filename, yamlData, 0644)
    if err != nil {
        panic("Unable to write data into the file")
    }
}

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Fatal(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Fatalf("%s is not a regular file", src)
	}

	source, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	defer source.Close()

	destination, err := os.Create(dst)

	if err != nil {
		log.Fatal(err)
	}

	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}