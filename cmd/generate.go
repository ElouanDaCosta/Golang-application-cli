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

type Folder struct {
	Name       string   `mapstructure:"name"`
	Subfolders []Folder `mapstructure:"subfolders"`
}

type Config struct {
	ServiceName string   `mapstructure:"service_name"`
	Port        int      `mapstructure:"port"`
	Folders     []Folder `mapstructure:"folders"`
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

	runGoModInit(config.ServiceName)

	createFolders(".", config.Folders)

	fmt.Printf("Microservice %s created successfully\n", config.ServiceName)
}

func runGoModInit(serviceName string) {
	cmd := exec.Command("go", "mod", "init", serviceName)
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to run go mod init: %v", err)
	}
}

func createFolders(basePath string, folders []Folder) {
	for _, folder := range folders {
		folderPath := fmt.Sprintf("%s/%s", basePath, folder.Name)
		os.Mkdir(folder.Name, 0755)
		createFolders(folderPath, folder.Subfolders)
	}
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

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("structure-file", "", "Pass the structure file name (default config.yaml)")
}
