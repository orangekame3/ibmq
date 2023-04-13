/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/orangekame3/ibmq/pkg"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
		device, _ := cmd.Flags().GetString("device")
		backendStatus(device)
	},
}

func init() {
	backendCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("device", "d", "", "device name eg. ibmq_qasm_simulator")
}

func backendStatus(device string) {
	token := pkg.LoadCredentials()
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
}
