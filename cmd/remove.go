package cmd

import (
	"fmt"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [doc name]",
	Short: "Remove an indexed document",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		docName := args[0]
		proc := internal.NewRAGProcessor(cfg)
		defer proc.MetaStore.Close()

		if err := proc.RemoveFile(docName); err != nil {
			fmt.Println("❌ Failed to remove:", err)
		}
	},
}
