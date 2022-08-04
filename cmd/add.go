package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add url",
	Short: "Add a podcast RSS feed",
	Long: `Add a podcast RSS feed using its URL.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := store.Add(args[0])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Added successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
