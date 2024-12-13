/*
Copyright © 2024 Fairblock
*/
package cmd

import (
	"fmt"
	"github.com/Fairblock/fairyport/config"
	"github.com/Fairblock/fairyport/internal/fairyport_app"
	"github.com/spf13/viper"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start fairyport relayer",
	Long:  `Start fairyport relayer`,

	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error loading config from file: %s\n", err.Error())
			return
		}

		viper.SetConfigName("config")
		viper.AddConfigPath(homeDir + "/.fairyport")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error loading config from file: %s\n", err.Error())
			return
		}

		err = viper.Unmarshal(&cfg)
		if err != nil {
			fmt.Printf("Error unmarshalling config from file: %s\n", err.Error())
			return
		}
		app := fairyport_app.NewFairyportApp(cfg)
		app.StartFairyport()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
