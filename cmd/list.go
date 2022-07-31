package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [channel]",
	Short: "Lists all channels, or if a channel is specified, episodes in a channel",
	Long: `Lists all channels, or if a channel is specified, episodes in a channel`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var list []string
		if len(args) == 0 {
			channels, err := store.ChannelList()
			cobra.CheckErr(err)
			list = channels
		} else {
			episodes, err := store.EpisodeList(args[0])
			cobra.CheckErr(err)
			list = episodes
		}
		for i, l := range list {
			fmt.Printf("%d: %s\n", i, l)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
