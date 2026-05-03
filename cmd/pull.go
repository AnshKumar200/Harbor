package cmd

import (
	"log"

	"github.com/AnshKumar200/Harbor/config"
	"github.com/AnshKumar200/Harbor/cmd/input"
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/AnshKumar200/Harbor/pkg/storage"
	"github.com/AnshKumar200/Harbor/pkg/util"
	"github.com/spf13/cobra"
)

var pullCommand = &cobra.Command{
	Use:   "pull IMAGE",
	Args:  cobra.ExactArgs(1),
	Short: "Pull container image",
	Run: func(cmd *cobra.Command, args []string) {
		regSvc := pkg.NewRegistryService(
			storage.NewImageStore(util.EnsureDir(config.DefaultImageStoreRootDir)),
		)

		name, tag := input.Parse(args[0])
		if err := regSvc.Pull(name, tag); err != nil {
			log.Fatalf("Failed to pull the image: %s\n", err)
		}
	},
}
