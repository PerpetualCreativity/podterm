package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/PerpetualCreativity/podterm/utils"
)

var store = utils.Store{ FeedName: "feed.xml" }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "podterm",
	Short: "podcast CLI and TUI client",
	Long: `podterm is an easy-to-use CLI and TUI client.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	d, err := os.UserCacheDir()
	cobra.CheckErr(err)
	rootFolder := filepath.Join(d, "podterm")
	_, err = os.Stat(rootFolder)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(rootFolder, 0750)
		cobra.CheckErr(err)
	}
	store.RootFolder = rootFolder
}


