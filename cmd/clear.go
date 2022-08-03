package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear channel OR clear --all",
	Short: "Delete all downloaded episodes from a channel, or all channels",
	Long: `Delete all downloaded episodes from a channel, or all channels`,
	Args: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		if (all&&len(args)!=0) || len(args)!=1 {
			return fmt.Errorf("must pass channel or --args, not both")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if all, _ := cmd.Flags().GetBool("all"); all {
			err := store.ClearAll()
			cobra.CheckErr(err)
		} else {
			chf, opts, err := store.FindChannel(args[0])
			cobra.CheckErr(err)
			if len(opts) > 1 {
				fmt.Printf("There are multiple matches for \"%s\".\n", args[0])
				fmt.Println("Since `clear` is a destructive command, the channel must be unambiguously specified.")
				fmt.Println("Matches:")
				for _, o := range opts {
					fmt.Printf("- %s\n", o)
				}
			} else {
				err = store.Clear(chf)
				cobra.CheckErr(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().Bool("all", false, "clear all channels")
}
