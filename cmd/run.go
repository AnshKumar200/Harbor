package cmd

import (
	"github.com/AnshKumar200/Harbor/cmd/input"
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/spf13/cobra"
)

var containerName string

var runCommand = &cobra.Command{
	Use:   "run IMAGE",
	Short: "Run a command in a new container",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName, imageTag := input.Parse(args[0])
		pkg.Run(imageName, imageTag, containerName)
	},
}

func init() {
	runCommand.Flags().StringVarP(&containerName, "name", "", "", "Assign a name to the container")
}
