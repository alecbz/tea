package cmd

import (
	"github.com/spf13/cobra"
)

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "Show the commit graph",
	Run: func(cmd *cobra.Command, args []string) {
		showGit("log",
			"--graph",
			"--format=%C(yellow)%h%C(reset)%C(auto)%d%C(reset)%n%C(blue)@%al%C(reset), %ar (%ad) %n%n%s%n",
			"--date", "format-local:%a, %b %d, %Y %I:%M %p",
		)
	},
}

func init() {
	rootCmd.AddCommand(commitsCmd)
}
