package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download channel I [J] [-o, --overwrite]",
	Short: "Download episode I, or all episodes from I to J",
	Long: `Downloads episode I, or all episodes from I to J
concurrently.
--overwrite causes podterm to overwrite an already downloaded
episode.`,
	Args: func(cmd *cobra.Command, args []string) error {
		l := len(args)
		if l<2 || l>3 {
			return fmt.Errorf("accepts between 1 and 2 args only")
		}
		I, err1 := strconv.Atoi(args[1])
		J := 0
		var err2 error
		if l==3 { J, err2 = strconv.Atoi(args[2]) }
		if err1 != nil || err2 != nil {
			return fmt.Errorf("args are not integers")
		}
		if I<0 || J<0 {
			return fmt.Errorf("args must be 0 or positive")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		I, _ := strconv.Atoi(args[0])
		J := I
		if len(args) == 2 {
			J, _ = strconv.Atoi(args[1])
		}
		var finish []chan bool
		for i:=I; i<J; i++ {
			done := make(chan bool)
			finish = append(finish, done)
			go func(i int) {
				_, err := store.GetEpisode(args[0], i, overwrite)
				if err != nil { fmt.Println(err) }
				done <- true
			}(i)
		}
		for _, done := range finish {
			<- done
			close(done)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().BoolP(
		"overwrite",
		"o",
		false,
		"download and overwrite even if already downloaded",
	)
}
