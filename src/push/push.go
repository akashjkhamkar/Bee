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

func Push(path, repository string) {
	config_file := path + "/config.yaml"
	configs := utils.Read_yaml_config_file(config_file)

	if !configs.Isbuilt {
		fmt.Println("Build the image first using 'bee build'")
		return
	}

	configs.Repository = repository

	tag(configs.Function_name, repository)
	push(repository)

	utils.Create_yaml_config_file(configs, config_file)
}