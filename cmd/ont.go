package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// ontCmd represents the ont command
var ontCmd = &cobra.Command{
	Use:   "ont",
	Short: "Open a new tab with the same path. Only works with Mac OS & iTerm2",
	Long: `
While working on the command-line, we always find ourselves in 
  situations where we have to open a new tab 
  and cd to the path we were working on.
  For example:Need another tab to run npm dev while
  making changes to code.

`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.Getwd()
		if err != nil {
			log.Panicln(err)
		}
		fmt.Print("Opening new tab: ")
		fmt.Println(path)
		cmd1 := exec.Command("osascript", "-e", `
		tell application "System Events" to tell process "Terminal" to keystroke "t" using command down
	`)

		if err := cmd1.Run(); err != nil {
			log.Panicln(err)
		}
	},
	Aliases: []string{"opennewtab"},
}

func init() {
	rootCmd.AddCommand(ontCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ontCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ontCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
