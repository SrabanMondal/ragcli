package cmd

import (
	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var cfg  = internal.DefaultConfig()

var rootCmd = &cobra.Command{
	Use:   "ragcli",
	Short: "RAG CLI using docai library",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cfg = internal.LoadConfig()

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(askCmd)
	rootCmd.AddCommand(summaryCmd)
	rootCmd.AddCommand(configCmd)
}
