package cmd

import (
	"fmt"
	"log"

	"github.com/AnshKumar200/Harbor/cmd/image"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCommand)
	rootCmd.AddCommand(runCommand)
	rootCmd.AddCommand(techCommand)
	rootCmd.AddCommand(image.Command)
}

var rootCmd = &cobra.Command{
	Use:   "harbor",
	Short: "harbor",
	Long:  "harbor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to harbor")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
