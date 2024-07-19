/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
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

type promptContent struct {
	errorMsg string
	label    string
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a microservice app",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("name")

		if appName != "" {
			generateFromStructureFile(appName)
		} else {
			generateFromStructureFile("new_app")
		}
	},
}

func generateFromStructureFile(appName string) {
	config := readStructureFile()
	newService := exec.Command("mkdir", appName)
	stdout, newServiceErr := newService.Output()

	if newServiceErr != nil {
		fmt.Println(newServiceErr.Error())
		return
	}

	fmt.Println(stdout)

	if err := os.Chdir(appName); err != nil {
		log.Fatalf("unable to change directory to %s, %v", appName, err)
	}

	runGoModInit(appName)

	createFolders(".", config.Folders)

	appType := promptContent{
		"Please select a package.",
		"Which package do you want your app to be based of ?",
	}

	newAppType := askUserForPackage(appType)
	addPackageToApp(newAppType, appName)

	fmt.Printf("Microservice %s created successfully\n", appName)
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
func readStructureFile() Config {
	viper.SetConfigName("config")
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

func askUserForPackage(pc promptContent) string {
	items := []string{"gin", "gRPC", "basic http"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.label,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func addPackageToApp(appType string, newAppBasePath string) {
	fmt.Println("user choose :", appType)
	if appType == "gin" {
		os.Chdir(newAppBasePath)
		exec.Command("go", "get", "-u", "github.com/gin-gonic/gin@latest").Output()
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().String("name", "", "Generate an app with the given name (default new_app)")
}
