package push

import (
	"bee/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func tag(function_name, repository string) {
	// Tag the image
	app := "docker"
	arg0 := "tag"
	arg1 := function_name
	arg2 := repository

	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func push(repository string) {
	// Push the image
	app := "docker"
	arg0 := "push"
	arg1 := repository

	cmd := exec.Command(app, arg0, arg1)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func create_service_file(function_name, path string) {
	service_config := utils.Yaml_service_config{
		APIVersion: "v1",
		Kind: "Service",
	}

	service_config.Metadata.Name = "bee-" + function_name + "-service"
	service_config.Spec.Selector.BeeFunction = function_name
	port_config := utils.Service_port{
		Protocol: "TCP",
		Port: 80,
		TargetPort: 8000,
	}
	service_config.Spec.Ports[0] = port_config

	utils.Create_yaml_config_file(service_config, path + "/Service.yaml")
}

func create_deployment_file(function_name, repository, path string, autoscaling bool, replicas int) {
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
	spec.Replicas = replicas
	spec.Selector.MatchLabels.BeeFunction = function_name
	spec.Template.Metadata.Labels.BeeFunction = function_name

	container_config := utils.Container_config{}
	container_config.Name = function_name
	container_config.Image = repository

	port_config := utils.Ports_config{}
	port_config.ContainerPort = 8000

	var ports [1] utils.Ports_config
	ports[0] = port_config
	container_config.Ports = ports

	var containers [1] utils.Container_config
	containers[0] = container_config

	spec.Template.Spec.Containers = containers
	deployment_config.Spec = spec

	utils.Create_yaml_config_file(deployment_config, path + "/Deployment.yaml")
}

func Push(path, repository string) {
	config_file := path + "/config.yaml"
	configs := utils.Read_yaml_config_file(config_file)

	if !configs.Isbuilt {
		fmt.Println("Build the image first using 'bee build'")
		return
	}

	configs.Repository = repository

	// tag and push
	tag(configs.Function_name, repository)
	push(repository)

	// make deployment.yaml and service.yaml
	// used in the deploy stage
	utils.Create_yaml_config_file(configs, config_file)
	create_deployment_file(configs.Function_name, repository, path, configs.Autoscaling, configs.Replicas)
	create_service_file(configs.Function_name, path)
}