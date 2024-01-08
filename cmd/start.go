/*
Copyright © 2024 Fairblock
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
	Short: "Start fairyport",
	Long:  `Start fairyport`,

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
}
