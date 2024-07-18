/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"service_name"`
	Port        int    `mapstructure:"port"`
	Structure   []struct {
		Name    string `mapstructure:"name"`
		Content []struct {
			Name string `mapstructure:"name"`
		}
	} `mapstructure:"structure"`
}

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
	config := readStructureFile(structureFile)
	newService := exec.Command("mkdir", config.ServiceName)
	stdout, newServiceErr := newService.Output()

	if newServiceErr != nil {
		fmt.Println(newServiceErr.Error())
		return
	}

	fmt.Println(stdout)

	if err := os.Chdir(config.ServiceName); err != nil {
		log.Fatalf("unable to change directory to %s, %v", config.ServiceName, err)
	}

	exec.Command("go", "mod", "init", config.ServiceName).Output()

	for i := range config.Structure {
		exec.Command("mkdir", config.Structure[i].Name).Output()
	}

	fmt.Printf("Microservice %s created successfully\n", config.ServiceName)
}

// pass the structure file to the flag without the extension
func readStructureFile(structureFile string) Config {
	viper.SetConfigName(structureFile)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return Config{}
	}
	fmt.Println(config)
	return config
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
	generateCmd.PersistentFlags().String("structure-file", "", "Pass the structure file name (default config.yaml)")
}
