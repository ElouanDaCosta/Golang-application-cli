/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"

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
	exec.Command("cat", "../storage/*")
}

func init() {
	rootCmd.AddCommand(listCmd)

}
