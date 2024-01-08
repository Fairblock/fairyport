/*
Copyright © 2024 Fairblock
*/
package cmd

import (
	"fmt"

	"github.com/Fairblock/fairyport/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize fairyport",
	Long:  `Initialize fairyport command creates default config for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		cfg := config.DefaultConfig()
		cfg.ExportConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
