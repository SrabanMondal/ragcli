package cmd

import (
	"fmt"
	"os"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [file path]",
	Short: "Add and index a PDF file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		// Ensure file exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println("❌ File not found:", file)
			return
		}

		proc := internal.NewRAGProcessor(cfg)
		defer proc.MetaStore.Close()

		if err := proc.AddFile(file); err != nil {
			fmt.Println("❌ Failed to add file:", err)
		}
	},
}
