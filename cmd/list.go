package cmd

import (
	"fmt"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all indexed documents",
	Run: func(cmd *cobra.Command, args []string) {
		proc := internal.NewRAGProcessor(cfg)
		defer proc.MetaStore.Close()

		if err := proc.ListFiles(); err != nil {
			fmt.Println("❌ Failed to list files:", err)
		}
	},
}
