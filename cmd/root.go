package cmd

import (
	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var cfg *internal.Config

var rootCmd = &cobra.Command{
	Use:   "ragcli",
	Short: "RAG CLI using docai library",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cfg = &internal.Config{
		DBPath:       "rag.db",
		EmbedderURL:  "http://localhost:11434/api/embeddings",
		GeneratorURL: "http://localhost:11434/api/generate",
		ChunkSize:    200,
		TopK:         3,
		Embed_Model: "nomic-embed-text",
		Generator_Model: "llama3.1",
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(askCmd)
}
