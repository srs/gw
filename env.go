package main

import (
	"os"
	"runtime"
	"strings"
)

func isUnix() bool {
	return !isWindows()
}

func isWindows() bool {
	return runtime.GOOS == `windows`
}

func getPathEnv() string {
	if isWindows() {
		return os.Getenv(`Path`)
	}

	return os.Getenv(`PATH`)
}

func getPaths() []string {
	env := getPathEnv()
	return strings.Split(env, string(os.PathListSeparator))
}

func getWorkingDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return pwd
}
