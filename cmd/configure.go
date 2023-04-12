/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure IBM Quantum API Token",
	Run: func(cmd *cobra.Command, args []string) {
		configure()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func configure() {
	fmt.Print("IBM Quantum API Token: ")

	apiTokenBytes, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("\nError reading API Token")
		os.Exit(1)
	}

	apiToken := string(apiTokenBytes)
	fmt.Println()

	ibmqDir := filepath.Join(os.Getenv("HOME"), ".ibmq")
	credsFile := filepath.Join(ibmqDir, "credentials")

	err = os.MkdirAll(ibmqDir, 0700)
	if err != nil {
		fmt.Printf("Error creating directory '%s': %v\n", ibmqDir, err)
		os.Exit(1)
	}

	err = os.WriteFile(credsFile, []byte("token="+apiToken), 0600)
	if err != nil {
		fmt.Printf("Error writing credentials file '%s': %v\n", credsFile, err)
		os.Exit(1)
	}

	fmt.Println("API Token saved")
}
