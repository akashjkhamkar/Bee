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

	// YAML file
	config := utils.Yaml_config{
		Function_name: function_name,
		Runtime: runtime,
		Replicas: 1,
	}

	utils.Create_yaml_config_file(config, function_name + "/config.yaml")

	// Tell user that template was successfully created
	fmt.Printf("Template successfully created inside folder '%s'\n", function_name)
}

var python_files = [] string {
	"requirements.txt",
	"python-function.py",
}