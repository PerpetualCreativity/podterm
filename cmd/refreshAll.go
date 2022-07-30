package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// refreshAllCmd represents the refreshAll command
var refreshAllCmd = &cobra.Command{
	Use:   "refresh-all [-n, --number N]",
	Short: "Downloads the last N (default: 5) episodes from all channels",
	Long: `Downloads the last N (default: 5) episodes from all channels`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("refreshAll called")
		l, _ := cmd.Flags().GetInt("number")
		err := store.RefreshAll(l)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(refreshAllCmd)

	refreshCmd.Flags().IntP("number", "n", 5, "number of episodes to download")
}
