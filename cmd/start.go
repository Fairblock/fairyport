/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"sync"

	"github.com/Fairblock/fairyport/app"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		wg.Add(1)

		withRelayer, _ := cmd.Flags().GetBool("with_relayer")

		if withRelayer {
			log.Println("Starting FairyPort with Hermes Relayer")

			// Call relayer's Run function concurrently
			relayerCmd.Run(cmd, []string{"start"})

			go func() {
				defer wg.Done()
				app := app.New()
				app.Start()
			}()
		} else {
			app := app.New()
			app.Start()
		}

		// Wait for both commands to finish
		wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().Bool("with_relayer", false, "Start hermes relayer")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
