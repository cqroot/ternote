package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/cqroot/ternote/pkg/ternote"
)

var (
	format  string
	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Dump database to yaml",
		Long:  "Dump database to yaml",
		Run:   RunDumpCmd,
	}
)

func init() {
	rootCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().StringVarP(&format, "format", "f", "yaml", "supported formats: yaml, markdown. default: yaml")
}

func RunDumpCmd(cmd *cobra.Command, args []string) {
	notes := ternote.New().Notes()

	switch format {
	case "markdown":
		result := ""

		prevCategory := ""

		for _, note := range notes {
			if note.Category != prevCategory {
				result = fmt.Sprintf(
					"%s- %s\n",
					result, note.Category,
				)
				prevCategory = note.Category
			}

			result = fmt.Sprintf(
				"%s  - [%s](notes/%s.md)\n",
				result, note.Title, note.Id,
			)
		}

		fmt.Print(result)

	default:
		b, err := yaml.Marshal(&notes)
		cobra.CheckErr(err)
		fmt.Println(string(b))
	}
}
