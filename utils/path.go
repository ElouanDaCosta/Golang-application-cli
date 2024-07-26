package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsolutePath() string {
	// Getting absolute path of current directory
	abs, err := filepath.Abs(".")

	// Printing if there is no error
	if err != nil {
		fmt.Println("Error getting the absolute path: ", err)
	}
	return abs
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
