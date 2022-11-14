package deploy

import (
	"bee/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func delete_existing_deployment(function_name string) {
	// Delete the exisiting deployment with the same name
	app := "kubectl"
	arg0 := "delete"
	arg1 := "deployment"
	arg2 := "bee-" + function_name

	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func apply_deployment(path, file string) {
	// Delete the exisiting deployment with the same name
	app := "kubectl"
	arg0 := "apply"
	arg1 := "-f"
	arg2 := path + "/" + file

	cmd := exec.Command(app, arg0, arg1, arg2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("Something went wrong while applying the deployment. ", err)
	}
}

func configure_autoscaling(config utils.Yaml_config) {
	if !config.Autoscaling {
		return
	}

	max := config.Max_replicas
	min := config.Min_replicas
	cpu := config.Cpu_percent

	if max <= 0 || min <= 0 || cpu <= 0 {
		log.Fatal("Mininum and Maximum number of replicas can't be 0 or lower.")
	}

	app := "kubectl"
	arg0 := "autoscale"
	arg1 := "deployment"
	arg2 := "bee-" + config.Function_name
	arg3 := "--cpu-percent=" + strconv.Itoa(cpu)
	arg4 := "--min=" + strconv.Itoa(min)
	arg5 := "--max=" + strconv.Itoa(max)

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("Something went wrong while applying the autoscaling configs. ", err)
	}
}

func Deploy(path string) {
	config := utils.Read_yaml_config_file(path + "/config.yaml")
	function_name := config.Function_name

	if !config.Isbuilt {
		fmt.Print("Build and push the function first by using 'bee build'")
		return
	}

	delete_existing_deployment(function_name)
	apply_deployment(path, "Deployment.yaml")
	apply_deployment(path, "Service.yaml")
	configure_autoscaling(config)
}