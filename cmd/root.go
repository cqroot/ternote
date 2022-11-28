package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cqroot/ternote/internal/tui"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ternote",
		Short: "Terminal note manager",
		Long:  "Terminal note manager",
		Run:   RunRootCmd,
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	cobra.CheckErr(tui.Run())
}
