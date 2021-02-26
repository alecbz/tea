package cmd

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View history of the current branch",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		branch, err := exec.Command("git", "rev-parse", "--abbrev-ref=strict", "HEAD").Output()
		if err != nil {
			log.Fatal(err)
		}

		reflog := exec.Command("git", "reflog", strings.TrimSpace(string(branch)))
		reflog.Stdout = os.Stdout
		if err := reflog.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
