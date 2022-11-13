package deploy

import (
	"bee/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func create_deployment_file(function_name, repository string, autoscaling bool, replicas int) {
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

	utils.Create_yaml_config_file(deployment_config, function_name + "/Deployment.yaml")
}

func delete_existing_deployment(function_name string) {
	// Delete the exisiting deployment with the same name
	app := "kubectl"
	arg0 := "delete"
	arg1 := "deployment"
	arg2 := "bee-" + function_name

	exec.Command(app, arg0, arg1, arg2)
}

func apply_deployment(path string) {
	// Delete the exisiting deployment with the same name
	app := "kubectl"
	arg0 := "apply"
	arg1 := "-f"
	arg2 := path + "/Deployment.yaml"

	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("Something went wrong while applying the deployment. ", err)
	}
}

func Deploy(path string) {
	config := utils.Read_yaml_config_file(path + "/config.yaml")
	function_name := config.Function_name
	replicas := config.Replicas
	repository := config.Repository
	autoscaling := config.Autoscaling

	if !config.Isbuilt {
		fmt.Print("Build and push the function first by using 'bee build'")
		return
	}

	create_deployment_file(function_name, repository, autoscaling, replicas)
	delete_existing_deployment(function_name)
	apply_deployment(path)
	// create a service for the deployment
}