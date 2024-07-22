/*
Copyright © 2024 Elouan DA COSTA PEIXOTO elouandacostapeixoto#gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// upgradeVersionCmd represents the upgradeVersion command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the go version of the specified application.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("name")

		if appName != "" {
			getAppPath(appName)
		} else {
			// getAppPath("new_app")
			getAllPath()
		}
	},
}

func getAllPath() []string {
	var outpout []string
	os.Chdir("./storage")
	f, err := os.ReadFile("app.txt")
	if err != nil {
		log.Println(err)
	}
	fileSplit := strings.Split(string(f), "\n")
	for _, value := range fileSplit {
		if strings.HasPrefix(value, "app path: ") {
			fmt.Println(value)
			outpout = append(outpout, value)
		}
	}
	return outpout
}

func getAppPath(appName string) string {
	os.Chdir("./storage")
	f, err := os.ReadFile("app.txt")
	if err != nil {
		log.Println(err)
	}
	outpout := strings.Split(string(f), "\n")
	for i := range outpout {
		if outpout[i] == "name: "+appName {
			fmt.Println(outpout[i+1])
			return outpout[i+1]
		}
	}
	return ""
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	upgradeCmd.PersistentFlags().String("name", "", "name of the app")
}
