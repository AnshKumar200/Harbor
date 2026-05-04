package cmd

import (
	"log"

	"github.com/AnshKumar200/Harbor/cmd/input"
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/spf13/cobra"
)

var techCommand = &cobra.Command{
	Use:   "tech",
	Short: "Technical subcommand used by harbor itself",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName, imageTag := input.Parse(args[0])

		conCmd := []string{}
		if len(args) >= 1 {
			conCmd = args[1:]
		}

		runtime := pkg.NewRuntimeService()
		if err := runtime.InitContainer(imageName, imageTag, containerName, conCmd); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	techCommand.Flags().StringVarP(&containerName, "name", "", "", "Assign a name to the container")
}
