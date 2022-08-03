package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh [channel]",
	Short: "Refresh all channel feeds, or only the specified channel feed.",
	Long: `Refresh all channel feeds, or only the specified channel feed.
Prints a list of all new episodes, grouped by channel.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		silent, _ := cmd.Flags().GetBool("silent")
		if len(args) == 0 {
			newCollections, err := store.RefreshAll()
			if err != nil { fmt.Println(err) }
			if !silent {
				for _, c := range newCollections {
					if len(c.Episodes) > 0 {
						fmt.Printf("%s: \n", c.Channel)
						for i, e := range c.Episodes {
							fmt.Printf("\t%d: %s\n", i, e.Title)
						}
					}
				}
			}
		} else {
			chf, _, err := store.FindChannel(args[0])
			cobra.CheckErr(err)
			newItems, err := store.Refresh(chf)
			if err != nil { fmt.Println(err) }
			if !silent {
				for i, e := range newItems {
					fmt.Printf("%d: %s", i, e.Title)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
	refreshCmd.Flags().BoolP("silent", "s", false, "suppress listing of new episodes")
}
