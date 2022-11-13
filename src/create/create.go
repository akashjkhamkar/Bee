package create

import (
	"bee/utils"
	"fmt"
	"log"
	"os"
)

var function_files = map[string] [] string {
	"python": python_files,
}

// TODO: in binary installation, the templates should be kept in the /bin or somewhere in the ~

func copy_function_files(files [] string, function_folder string) {
	for _, file := range files {
		src := "templates/python/" + file
		utils.Copy(src, function_folder + "/" + file)
	}
}

func Create(runtime , function_name string) {
	files, is_present := function_files[runtime]
	
	if !is_present {
		log.Fatalf("%s : No such runtime", runtime)
	}
	
	// The folder with name
	if err := os.Mkdir(function_name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Copy the skeleton code into the folder
	copy_function_files(files, function_name)

	// YAML configs file
	config := utils.Yaml_config{
		Function_name: function_name,
		Runtime: runtime,
		Replicas: 1,
	}

	utils.Create_yaml_config_file(config, function_name + "/config.yaml")
	
	// Deployment file
	deployment_config := utils.Yaml_deployment{}

	// apu
	deployment_config.APIVersion = "apps/v1"
	deployment_config.Kind = "Deployment"

	// metadata
	metadata := deployment_config.Metadata
	metadata.Name = "bee-" + function_name
	metadata.Labels.BeeFunction = function_name
	deployment_config.Metadata = metadata

	// spec
	spec := deployment_config.Spec
	spec.Replicas = 1
	spec.Selector.MatchLabels.BeeFunction = function_name
	spec.Template.Metadata.Labels.App = "bee-" + function_name

	container_config := utils.Container_config{}
	container_config.Name = function_name

	port_config := utils.Ports_config{}
	port_config.ContainerPort = 8000

	var ports [1] utils.Ports_config
	ports[0] = port_config
	container_config.Ports = ports

	var containers [1] utils.Container_config
	containers[0] = container_config

	spec.Template.Spec.Containers = containers
	deployment_config.Spec = spec

	utils.Create_yaml_config_file(deployment_config, function_name + "/Deployment.yaml")

	fmt.Printf("Template successfully created inside folder '%s'\n", function_name)
}

var python_files = [] string {
	"requirements.txt",
	"function.py",
	"flask-listener.py",
	"Dockerfile",
}