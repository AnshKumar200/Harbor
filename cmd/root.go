package cmd

import (
	"log"

	"github.com/AnshKumar200/Harbor/cmd/image"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCommand)
	rootCmd.AddCommand(runCommand)
	rootCmd.AddCommand(image.Command)
	rootCmd.AddCommand(rmCommand)
	rootCmd.AddCommand(psCommand)
	rootCmd.AddCommand(internalCommand)
}

var rootCmd = &cobra.Command{
	Use:   "harbor",
	Short: "harbor",
	Long:  "harbor",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
