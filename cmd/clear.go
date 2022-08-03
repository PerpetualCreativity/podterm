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
		var err error
		if all, _ := cmd.Flags().GetBool("all"); all {
			err = store.ClearAll()
		} else {
			err = store.Clear(args[0])
		}
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().Bool("all", false, "clear all channels")
}
