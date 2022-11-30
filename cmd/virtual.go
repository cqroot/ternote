package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cqroot/ternote/pkg/ternote"
	"github.com/cqroot/ternote/pkg/virtual"
)

var (
	virtualCmd = &cobra.Command{
		Use:   "virtual",
		Short: "",
		Long:  "",
		Run:   RunVirtualCmd,
	}
)

func init() {
	rootCmd.AddCommand(virtualCmd)
}

func RunVirtualCmd(cmd *cobra.Command, args []string) {
	for _, note := range ternote.New().Notes() {
		cobra.CheckErr(virtual.NewNote(note))

		fmt.Printf("%s: [ %s ] - %s\n", note.Id, note.Category, note.Title)
	}
}
