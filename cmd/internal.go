package cmd

import (
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var internReq pkg.RunRequest

var internalCommand = &cobra.Command{
	Use:   "internal",
	Short: "Internal command for harbor itself",
	Run: func(cmd *cobra.Command, args []string) {
		r := pkg.NewRuntimeService()
		if err := r.InitContainer(internReq); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	internalCommand.Flags().StringVarP(&internReq.ContainerName, "ContainerName", "", "", "")
	internalCommand.Flags().StringVarP(&internReq.ImageName, "ImageName", "", "", "")
	internalCommand.Flags().StringVarP(&internReq.ImageTag, "ImageTag", "", "", "")
	internalCommand.Flags().StringVarP(&internReq.ContainerCommand, "ContainerCommand", "", "", "")
	internalCommand.Flags().StringVarP(&internReq.ContainerID, "ContainerID", "", "", "")
}
