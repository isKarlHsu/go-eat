package command

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "root",
	Short: "根命令",
	Long: "根命令",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("command execute success")
	},
}

func Execute()  {
	rootCmd.Execute()
}