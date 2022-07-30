package cmd

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove channel",
	Short: "Remove a channel by name",
	Long: `Remove a channel by name. All downloaded episodes will be deleted.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := store.Remove(args[0])
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
