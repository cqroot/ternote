package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ternote/pkg/config"
)

var (
	pathCmd = &cobra.Command{
		Use:   "path",
		Short: "Print application data path",
		Long:  "Print application data path",
		Run:   RunPathCmd,
	}
)

func init() {
	rootCmd.AddCommand(pathCmd)
}

func RunPathCmd(cmd *cobra.Command, args []string) {
	basePath, err := config.BasePath()
	cobra.CheckErr(err)

	fmt.Println(basePath)
}
