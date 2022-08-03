package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh [channel]",
	Short: "Refresh all channel feeds, or specified channel feed only.",
	Long: `Refresh all channel feeds, or specified channel feed only.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			err = store.RefreshAll()
		} else {
			err = store.Refresh(args[0])
		}
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
