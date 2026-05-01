package cmd

import (
	"github.com/AnshKumar200/Harbor/pkg"
	"github.com/spf13/cobra"
)

var techCommand = &cobra.Command{
	Use:   "tech",
	Short: "technical subcommand used by harbor itself",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Chroot()
	},
}
