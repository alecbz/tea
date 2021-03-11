package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Bring the branch up-to-date with the main branch",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		spin("Rebasing", func() {
			runGit("pull", "--rebase", "origin", config.MainBranch)
		})
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
