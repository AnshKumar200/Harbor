package cmd

import (
	"log"

	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/spf13/cobra"
)

var techCommand = &cobra.Command{
	Use:   "tech",
	Short: "Technical subcommand used by harbor itself",
	Run: func(cmd *cobra.Command, args []string) {
		runtime := pkg.NewRuntimeService()
		if err := runtime.InitContainer(args[0], args[1], args[2]); err != nil {
			log.Fatal(err)
		}
	},
}
