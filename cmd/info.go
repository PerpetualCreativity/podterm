package cmd

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"path/filepath"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PerpetualCreativity/podterm/utils"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info channel [episode_index]",
	Short: "Get detailed information about an episode",
	Long: `Get detailed information about an episode`,
	Args: cobra.MatchAll(
		cobra.RangeArgs(1, 2),
		func(cmd *cobra.Command, args []string) error {
			if len(args) == 2 {
				i, err := strconv.Atoi(args[1])
				if err != nil { return fmt.Errorf("episode_index must be an integer") }
				if i<0 { return fmt.Errorf("episode_index must be nonnegative") }
			}
			return nil
		},
	),
	Run: func(cmd *cobra.Command, args []string) {
		ch, err := utils.ParseFile(filepath.Join(store.RootFolder, args[0], store.FeedName))
		cobra.CheckErr(err)
		info := ""
		if len(args) == 1 {
			info = fmt.Sprintf(
				"# %s\n\n(%s)\n\n%s\n\nCopyright %s\n",
				ch.Title,
				ch.Link,
				ch.Description,
				ch.Copyright,
			)
		} else {
			i, _ := strconv.Atoi(args[1])
			episode := ch.Items[i]
			var time string
			ts := strings.Split(episode.Duration, ":")
			if len(ts) == 2 {
				time = fmt.Sprintf("%sm%ss", ts[0], ts[1])
			} else if len(ts) == 3 {
				time = fmt.Sprintf("%sh%sm%ss", ts[0], ts[1], ts[2])
			} else {
				time = episode.Duration
			}
			desc, _ := md.
				NewConverter("", true, nil).
				ConvertString(episode.Description)
			info = fmt.Sprintf(
				"# %s\n\nThis episode is %s long and was published on %s\n\n%s\n",
				episode.Title,
				time,
				episode.PubDate,
				desc,
			)
		}
		r, _ := glamour.NewTermRenderer(glamour.WithAutoStyle())
		out, _ := r.Render(info)
		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}