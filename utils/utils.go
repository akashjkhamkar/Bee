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
	Repository string;
	Replicas int;
	Autoscaling bool;
	Isbuilt bool;
}

type Yaml_deployment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			BeeFunction string `yaml:"bee-function"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Replicas int `yaml:"replicas"`
		Selector struct {
			MatchLabels struct {
				BeeFunction string `yaml:"bee-function"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Labels struct {
					BeeFunction string `yaml:"bee-function"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Containers [1] Container_config
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

type Container_config struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Ports [1] Ports_config `yaml:"ports"`
}

type Ports_config struct {
	ContainerPort int `yaml:"containerPort"`
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

func Read_yaml_deployment_file(filename string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Error reading YAML file: %s\n", err)
    }

	err = yaml.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		log.Fatalf("error while reading the Deployments file: %v", err)
	}

	return m
}

func Create_yaml_config_file(data interface{}, filename string) {
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