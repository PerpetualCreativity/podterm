package cmd

import (
	"fmt"
	"github.com/PerpetualCreativity/podterm/utils"
	"github.com/spf13/cobra"
	"path/filepath"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [channel] [-l, --length N] [-d, --downloaded-only]",
	Short: "Lists all channels, or if a channel is specified, episodes in a channel",
	Long: `Lists all channels, or if a channel is specified, episodes in a channel.
Only lists latest 10 episodes by default; this can be changed with --length. If N<0,
all episodes will be listed (piping to a pager is encouraged).
If --downloaded-only is passed, only downloaded episodes are listed.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			channels, err := store.ChannelList()
			cobra.CheckErr(err)
			for _, l := range channels {
				fmt.Printf("- %s\n", l)
			}
		} else {
			length, err := cmd.Flags().GetInt("length")
			downloaded, _ := cmd.Flags().GetBool("downloaded-only")
			chf, _, err := store.FindChannel(args[0])
			cobra.CheckErr(err)
			var items []utils.Item
			if downloaded {
				i, err := store.DownloadedEpisodeList(chf)
				cobra.CheckErr(err)
				items = i
			} else {
				ch, err := utils.ParseFile(filepath.Join(store.RootFolder, chf, store.FeedName))
				cobra.CheckErr(err)
				items = ch.Items
			}
			b := length
			if length<0 || len(items)<length {
				b = len(items)
			}
			for i := 0; i<b; i++ {
				fmt.Printf("#%d: %s\n", i, items[i].Title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP(
		"downloaded-only",
		"d",
		false,
		"list only downloaded episodes",
	)
	listCmd.Flags().IntP(
		"length",
		"l",
		10,
		"maximum number of episodes to list",
	)
}
