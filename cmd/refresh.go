package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh title [-n,--number N] [-o,--overwrite]",
	Short: "Download latest N episodes from channel",
	Long: `Refresh the channel specified by title, by getting
the new feed, then downloading the last N episodes from that channel.
If N is 0, only the feed is refreshed. By default, already downloaded
episodes are skipped (--overwrite removes this check).`,
	Args: cobra.RangeArgs(1,2),
	Run: func(cmd *cobra.Command, args []string) {
		l, _ := cmd.Flags().GetInt("number")
		ow, _ := cmd.Flags().GetBool("overwrite")
		err := store.Refresh(args[0], l, ow)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Refreshed last %d episodes of %s successfully\n", l, args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
	refreshCmd.Flags().IntP("number", "n", 5, "number of episodes to download")
	refreshCmd.Flags().BoolP("overwrite", "o", false, "overwrite already downloaded episodes")
}
