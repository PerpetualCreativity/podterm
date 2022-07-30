package cmd

import (
	"github.com/spf13/cobra"
)

// refreshAllCmd represents the refreshAll command
var refreshAllCmd = &cobra.Command{
	Use:   "refresh-all [-n, --number N]",
	Short: "Downloads the last N (default: 5) episodes from all channels",
	Long: `Download the last N (default: 5) episodes from all channels`,
	Run: func(cmd *cobra.Command, args []string) {
		l, _ := cmd.Flags().GetInt("number")
		err := store.RefreshAll(l)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(refreshAllCmd)

	refreshAllCmd.Flags().IntP("number", "n", 5, "number of episodes to download")
}
