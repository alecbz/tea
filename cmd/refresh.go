package cmd

import (
	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
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
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
