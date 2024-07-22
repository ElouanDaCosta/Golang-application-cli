/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List generated app.",
	Long:  `List all generated app or one app by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("name")

		if appName != "" {
			getOneApp(appName)
		} else {
			listAllApp()
		}
	},
}

func listAllApp() {
	os.Chdir("./storage")
	f, err := os.ReadFile("app.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(f))
}

func getOneApp(appName string) {
	os.Chdir("./storage")
	f, err := os.ReadFile("app.txt")
	if err != nil {
		log.Println(err)
	}
	outpout := strings.Split(string(f), "\n")
	for i := range outpout {
		if outpout[i] == "name: "+appName {
			fmt.Println(outpout[i])
			fmt.Println(outpout[i+1])
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().String("name", "", "Return the app with the given name (default new_app)")
}
