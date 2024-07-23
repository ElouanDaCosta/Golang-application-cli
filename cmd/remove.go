/*
Copyright © 2024 Elouan DA COSTA PEIXOTO elouandacostapeixoto@gmail.com
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an application from the storage",
	Long: `Remove an application from the storage or remove an application with the rf flag. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		appName, _ := cmd.Flags().GetString("name")
		if appName != "" {
			removeFromStorage(appName)
		} else {
			listAllApp()
		}
	},
}

func removeFromStorage(appName string) {
	os.Chdir("./storage")

	filePath := "app.txt"
	targetLine := appName

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	index := 0
	for i, line := range lines {
		if line == "name: "+targetLine {
			index = i
			break
		}
	}

	if index == 0 {
		fmt.Println("Line not found")
		return
	}

	var newLines []string
	if index < len(lines)-1 {
		newLines = append(lines[:index], lines[index+2:]...)
	} else {
		newLines = lines[:index]
	}

	output := strings.Join(newLines, "\n")
	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalf("unable to write file: %v", err)
	}

	fmt.Println("Lines deleted successfully")
}

func cleanAllStorage() {
	os.Chdir("./storage")

	f, err := os.OpenFile("app.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var bs []byte
	buf := bytes.NewBuffer(bs)

	var text string
	for scanner.Scan() {
		text = scanner.Text()
		_, err := buf.WriteString(text + "\n")
		if err != nil {
			panic("Couldn't replace line")
		}
	}
	f.Truncate(0)
	f.Seek(0, 0)
	buf.WriteTo(f)
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.PersistentFlags().String("name", "", "Return the app with the given name (default new_app)")
}
