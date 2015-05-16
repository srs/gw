package main

import (
	"os"
	"os/exec"
)

func main() {
	name := findGradle()

	cmd := exec.Command(name, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
