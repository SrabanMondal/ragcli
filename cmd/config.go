package cmd

import (
	"fmt"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure runtime settings like models, chunk size, DB path, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.DBPath, _ = cmd.Flags().GetString("db")
		cfg.EmbedderURL, _ = cmd.Flags().GetString("embed-url")
		cfg.GeneratorURL, _ = cmd.Flags().GetString("gen-url")
		cfg.Embed_Model, _ = cmd.Flags().GetString("embed-model")
		cfg.Generator_Model, _ = cmd.Flags().GetString("gen-model")
		cfg.ChunkSize, _ = cmd.Flags().GetInt("chunk")
		cfg.TopK, _ = cmd.Flags().GetInt("topk")

		fmt.Println("✅ Configuration set:")
		fmt.Printf("🗂️ DB: %s\n", cfg.DBPath)
		fmt.Printf("🔗 Embedder URL: %s\n", cfg.EmbedderURL)
		fmt.Printf("🧠 Generator URL: %s\n", cfg.GeneratorURL)
		fmt.Printf("📦 Embed Model: %s\n", cfg.Embed_Model)
		fmt.Printf("🤖 Gen Model: %s\n", cfg.Generator_Model)
		fmt.Printf("✂️ Chunk Size: %d\n", cfg.ChunkSize)
		fmt.Printf("🎯 Top-K: %d\n", cfg.TopK)
		internal.SaveConfig(cfg)
	},
}

func init() {
	configCmd.Flags().String("db", cfg.DBPath, "Path to SQLite DB file")
	configCmd.Flags().String("embed-url", cfg.EmbedderURL, "Embedder API endpoint")
	configCmd.Flags().String("gen-url", cfg.GeneratorURL, "Generator API endpoint")
	configCmd.Flags().String("embed-model", cfg.Embed_Model, "Embedding model name")
	configCmd.Flags().String("gen-model", cfg.Generator_Model, "Generation model name")
	configCmd.Flags().Int("chunk", cfg.ChunkSize, "Chunk size for document splitting")
	configCmd.Flags().Int("topk", cfg.TopK, "Number of top chunks to retrieve")
}
