package cmd

import (
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
		chf, _, err := store.FindChannel(args[0])
		cobra.CheckErr(err)
		i := 0
		if len(args) == 2 {
			in, err := strconv.Atoi(args[1])
			cobra.CheckErr(err)
			i = in
		}
		path, err := store.GetEpisode(chf, i, false)
		cobra.CheckErr(err)
		_, err = exec.Command("open", path).Output()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
