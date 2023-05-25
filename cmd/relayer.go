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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// relayerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// relayerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
