package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetGoVersion(appPath string) (string, error) {
	os.Chdir(appPath)
	f, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	lines := strings.Split(string(f), "\n")

	versionSplit := strings.Split(lines[2], "go")

	return versionSplit[1], nil
}
