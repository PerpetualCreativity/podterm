package cmd

import (
	"fmt"
	"github.com/PerpetualCreativity/podterm/utils"
	"github.com/spf13/cobra"
	"path/filepath"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [channel]",
	Short: "Lists all channels, or if a channel is specified, episodes in a channel",
	Long: `Lists all channels, or if a channel is specified, episodes in a channel`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			channels, err := store.ChannelList()
			cobra.CheckErr(err)
			for i, l := range channels {
				fmt.Printf("%d: %s\n", i, l)
			}
		} else {
			ch, err := utils.ParseFile(filepath.Join(store.RootFolder, args[0], store.FeedName))
			cobra.CheckErr(err)
			for i, episode := range ch.Items {
				fmt.Printf("#%d: %s\n", i, episode.Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
