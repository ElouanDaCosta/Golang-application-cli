/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate microservice structure",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		structureFile, _ := cmd.Flags().GetString("structure-file")

		if structureFile != "" {
			generateFromStructureFile(structureFile)
		} else {
			generateFromUserInput()
		}
	},
}

func generateFromStructureFile(structureFile string) {
	fmt.Println(structureFile)
}

func generateFromUserInput() {
	fmt.Println("input text:")
	var w1 string
	n, err := fmt.Scanln(&w1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read line: %s-\n", w1)

	args0 := w1

	cmd := exec.Command("mkdir", args0)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("structure-file", "", "Pass the structure file name and generate project structure from it.")
}
