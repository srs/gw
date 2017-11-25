package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mgutz/ansi"
)

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func getGradleExec() string {
	if isWindows() {
		return "gradle.bat"
	}

	return "gradle"
}

func findGradleExec() string {
	exec := getGradleExec()
	paths := getPaths()

	for i := range paths {
		name := filepath.Join(paths[i], exec)
		if fileExists(name) {
			return name
		}
	}

	printNoGradleError(exec)
	os.Exit(-1)
	return exec
}

func getGradleWExec() string {
	if isWindows() {
		return "gradlew.bat"
	}

	return "gradlew"
}

func findGradleWExec(dir string) (string, error) {
	exec := getGradleWExec()
	parent := filepath.Join(dir, "..")

	if parent == dir {
		return exec, errors.New("gradle wrapper not found")
	}

	name := filepath.Join(dir, exec)
	if fileExists(name) {
		return name, nil
	}

	return findGradleWExec(parent)
}

func findGradle() string {
	pwd := getWorkingDir()

	exec, err := findGradleWExec(pwd)
	if err != nil {
		printNoGradleWNotice(exec)
		return findGradleExec()
	}

	return exec
}

func printNoGradleWNotice(exec string) {
	fmt.Println()
	fmt.Println(ansi.Color("NOTICE:", "yellow+b"))
	fmt.Printf("No %s set up for this project. ", exec)
	fmt.Println("Please consider setting one up.")
	fmt.Println("(http://gradle.org/docs/current/userguide/gradle_wrapper.html)")
	fmt.Println()
}

func printNoGradleError(exec string) {
	fmt.Println()
	fmt.Println(ansi.Color("ERROR:", "red+b"))
	fmt.Println("No %s exec found in path. Please install gradle.")
	fmt.Println("(http://gradle.org/docs/current/userguide/installation.html)")
	fmt.Println()
}
