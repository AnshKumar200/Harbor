package image

import (
	"fmt"
	"log"

	"github.com/AnshKumar200/Harbor/cmd/input"
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "rm IMAGE",
	Args:  cobra.ExactArgs(1),
	Short: "Remove image",
	Run: func(cmd *cobra.Command, args []string) {
		imageName, imageTag := input.Parse(args[0])
		runtime := pkg.NewRuntimeService()
		if err := runtime.RemoveImage(imageName, imageTag); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Image deleted: %s:%s\n", imageName, imageTag)
		}
	},
}
