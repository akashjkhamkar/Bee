package push

import (
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

func Push(function_name, repository string) {
	tag(function_name, repository)
	push(repository)
}