/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/orangekame3/ibmq/model"
	"github.com/orangekame3/ibmq/pkg"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ls called")
		listBackends()
	},
}

var longOutput bool

func init() {
	backendCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&longOutput, "long", "l", false, "Display detailed information about backends")
}

func listBackends() {
	token := pkg.LoadCredentials()
	resp, err := pkg.GetRequest(token, "backends")
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}
	var backendList model.BackendList
	err = json.Unmarshal(body, &backendList)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	for _, device := range backendList.Devices {
		if longOutput {
			backendDetails, err := pkg.GetBackendStatus(token, device)
			if err != nil {
				fmt.Println("Error creating request:", err)
				os.Exit(1)
			}
			fmt.Printf("- %s\n", device)
			fmt.Printf("  Backend version: %s\n", backendDetails.BackendVersion)
			fmt.Printf("  State: %v\n", backendDetails.State)
			fmt.Printf("  Message: %s\n", backendDetails.Message)
			fmt.Printf("  Length Queue: %d\n", backendDetails.LengthQueue)
			fmt.Println()
		} else {
			fmt.Printf("- %s\n", device)
		}
	}
}
