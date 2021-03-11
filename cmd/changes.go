package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// changesCmd represents the changes command
var changesCmd = &cobra.Command{
	Use:   "changes",
	Short: "Changes on this branch compared to main",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showGit("diff", fmt.Sprintf("origin/%s...", config.MainBranch))
	},
}

func init() {
	rootCmd.AddCommand(changesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
