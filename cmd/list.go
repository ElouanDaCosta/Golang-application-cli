/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of generated app.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listAllApp()
	},
}

func listAllApp() {
	os.Chdir("./storage")
	f, _ := os.ReadFile("app.txt")
	getAppData(f)
}

func getAppData(appFile []byte) []string {
	appData := string(appFile)
	var outpout []string
	appDataArray := strings.Split(appData, "\n")
	name := strings.Split(appDataArray[0], "name: ")
	path := strings.Split(appDataArray[1], "app path: ")
	outpout = append(outpout, name[1], path[1])
	return outpout
}

func init() {
	rootCmd.AddCommand(listCmd)

}
