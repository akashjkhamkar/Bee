package create

import (
	"bee/utils"
	"log"
	"os"
)

var function_files = map[string] [] string {
	"python": python_files,
}

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
	// the name
	// runtime
	// auto-scaling
	// replicas: 1
	// use some yaml generator to generate the file
}


var python_files = [] string {
	"requirements.txt",
	"python-function.py",
}