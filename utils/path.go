package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(filePath string) string {
	// Getting absolute path of hello.go
	abs, err := filepath.Abs(filePath)

	// Printing if there is no error
	if err != nil {
		fmt.Println("Error getting the absolute path: ", err)
	}
	return abs
}

func RunProgram(program string) string {
	a := strings.Split(program, " ")
	out, err := exec.Command(a[0], a[1:]...).Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func GetInstalledPath() string {
	// getting the environment variable of the installed path of the cli
	getPathEnv, exists := os.LookupEnv("GAC")
	if !exists {
		fmt.Println("Error reading GAC variables")
		return ""
	}
	return getPathEnv
}
