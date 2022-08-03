package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove channel",
	Short: "Remove a channel by name",
	Long: `Remove a channel by name. All downloaded episodes will be deleted.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		chf, opts, err := store.FindChannel(args[0])
		cobra.CheckErr(err)
		if len(opts) > 1 {
			fmt.Printf("There are multiple matches for \"%s\".\n", args[0])
			fmt.Println("Since `remove` is a destructive command, the channel must be unambiguously specified.")
			fmt.Println("Matches:")
			for _, o := range opts {
				fmt.Printf("- %s\n", o)
			}
		} else {
			err = store.Remove(chf)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
