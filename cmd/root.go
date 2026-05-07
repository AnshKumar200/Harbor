package cmd

import (
	"github.com/AnshKumar200/Harbor/cmd/image"
	"github.com/AnshKumar200/Harbor/config"
	"github.com/AnshKumar200/Harbor/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbosity int

func init() {
	rootCmd.AddCommand(pullCommand)
	rootCmd.AddCommand(runCommand)
	rootCmd.AddCommand(image.Command)
	rootCmd.AddCommand(rmCommand)
	rootCmd.AddCommand(psCommand)
	rootCmd.AddCommand(internalCommand)

	rootCmd.PersistentFlags().IntVarP(&verbosity, "verbose", "v", int(config.DefaultLogLevel), "0 to 6: Trace, Debug, Info, Warn, Error, Fatal, Panic")
	logrus.SetLevel(logging.VerbosityToLogrusLevel(verbosity))
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
		logrus.Fatal(err)
	}
}
