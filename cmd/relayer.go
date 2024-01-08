package cmd

import (
	"fmt"
	"io"

	"os/exec"

	"github.com/spf13/cobra"
)

// relayerCmd represents the relayer command
var relayerCmd = &cobra.Command{
	Use:   "relayer",
	Short: "Start relayer",
	Long:  `Start hermes relayer for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		wg.Add(1)
		defer wg.Done()
		cmdHermes := exec.Command("hermes", args...)

		stdout, _ := cmdHermes.StdoutPipe()
		stderr, _ := cmdHermes.StderrPipe()

		err := cmdHermes.Start()
		if err != nil {
			fmt.Println("Error starting binary:", err)
			return
		}

		outputBytes, _ := io.ReadAll(stdout)
		errorBytes, _ := io.ReadAll(stderr)

		output := string(outputBytes)
		errorOutput := string(errorBytes)

		fmt.Println(output)
		fmt.Println(errorOutput)

		err = cmdHermes.Wait()
		if err != nil {
			fmt.Println("Error waiting for binary:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(relayerCmd)
}
