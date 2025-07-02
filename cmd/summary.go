package cmd

import (
	"fmt"

	"github.com/SrabanMondal/ragcli/internal"
	"github.com/spf13/cobra"
)

var path string
var summaryCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize any document",
	Run: func(cmd *cobra.Command, args []string) {
		proc := internal.NewRAGProcessor(cfg)
		if path==""{
			fmt.Println("❌ Please provide a document path")
			return
		}
		if err := proc.Summarize(path); err != nil {
			fmt.Println("❌ Failed to summarize:", err)
		}
	},
}

func init(){
	summaryCmd.Flags().StringVar(&path,"path","","Provide file path to summarize")
}
