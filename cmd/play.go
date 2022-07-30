package cmd

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play channel [n]",
	Short: "Play an episode from a channel",
	Long: `Play the latest episode from the channel, or the
Nth episode in reverse-chronological order.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		i := 0
		if len(args) == 2 {
			in, err := strconv.Atoi(args[1])
			cobra.CheckErr(err)
			i = in
		}
		path, err := store.GetEpisodePath(args[0], i)
		cobra.CheckErr(err)
		_, err = exec.Command(fmt.Sprintf("open %s", path)).Output()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
