package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadCredentials() string {
	credsFile := filepath.Join(os.Getenv("HOME"), ".ibmq", "credentials")
	viper.SetConfigFile(credsFile)
	viper.SetConfigType("properties")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading credentials file: %v\n", err)
		os.Exit(1)
	}

	return viper.GetString("token")
}
