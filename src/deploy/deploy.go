package deploy

import (
	"bee/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func delete_existing_deployment(function_name string) {
	// Delete the exisiting deployment with the same name
	app := "kubectl"
	arg0 := "delete"
	arg1 := "deployment"
	arg2 := "bee-" + function_name

	exec.Command(app, arg0, arg1, arg2)
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
}