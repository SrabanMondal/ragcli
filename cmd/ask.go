package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var docName string
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question from your indexed documents",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("❓ Enter your question: ")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		if query == "" {
			fmt.Println("⚠️ Empty question. Try again.")
			return
		}

		proc := internal.NewRAGProcessor(cfg)
		defer proc.MetaStore.Close()

		answer, err := proc.AskQuestion(query, docName)
		if err != nil {
			fmt.Println("❌ Failed to generate answer:", err)
			return
		}

		fmt.Println("\n🧠 Answer:\n", answer)
	},
}

func init(){
	askCmd.Flags().StringVarP(&docName, "document","d","","Add specifc document name to ask question")
}