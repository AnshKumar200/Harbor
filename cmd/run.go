package cmd

import (
	"github.com/AnshKumar200/Harbor/cmd/input"
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var req pkg.RunRequest

var runCommand = &cobra.Command{
	Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Run a command in a new container",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageName, imageTag := input.Parse(args[0])
		req.ImageName = imageName
		req.ImageTag = imageTag
		if len(args) > 1 {
			req.ContainerCommand = args[1]
		}
		if len(args) > 2 {
			req.ContainerArgs = args[2:]
		}

		runtime := pkg.NewRuntimeService()

		if err := runtime.Run(req); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	runCommand.Flags().StringVarP(&req.ContainerName, "name", "", "", "Assign a name to the container")
}
