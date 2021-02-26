package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [branch-name]",
	Short: "Start a new branch",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("git", "fetch", "origin", "master-passing-tests").Run(); err != nil {
			log.Fatal("error fetching:", err)
		}

		u, err := user.Current()
		if err != nil {
			log.Fatal("error getting current user:", err)
		}
		if err := exec.Command("git", "switch", "--create", fmt.Sprintf("%s/%s", u.Username, args[0]), "origin/master-passing-tests").Run(); err != nil {
			log.Fatal("error starting new branch:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
