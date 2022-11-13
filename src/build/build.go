package build

import (
	"bee/utils"
	"log"
	"os"
	"os/exec"
)

// TODO: some data cleaning in cmd

func Build(path string) {
	config_file := path + "/config.yaml"
	configs := utils.Read_yaml_config_file(config_file)

	app := "docker"
    arg0 := "build"
    arg1 := "-t"
    arg2 := configs.Function_name + ":latest"
    arg3 := path

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }

	configs.Isbuilt = true
	utils.Create_yaml_config_file(configs, config_file)
}