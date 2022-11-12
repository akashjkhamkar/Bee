package create

import (
	"fmt"
	"log"
	"os"
)

var creators = map[string] func() {
	"python": create_python,
}

func Create(runtime , filename string) {
	// The folder with name
	if err := os.Mkdir(filename, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	creator, is_present := creators[runtime]
	
	if !is_present {
		log.Fatalf("No such runtime %s", runtime)
	}

	// Copy the skeleton code
	creator()

	// YAML file
	// the name
	// runtime
	// auto-scaling
	// replicas: 1
	// use some yaml generator to generate the file
}


func create_python() {
	fmt.Println("Creating a python function")
}