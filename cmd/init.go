package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func run(name string, args ...string) {
	if err := exec.Command(name, args...).Run(); err != nil {
		panic(err)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		run("git", "init", "--initial-branch", "master", cwd)

		run("git", "commit", "--allow-empty", "--message", "root commit")

		// Set up a default "origin" remote as our own repo.
		run("git", "remote", "add", "origin", fmt.Sprintf("file://%s", cwd))
		run("git", "fetch", "origin", "master")

		run("git", "checkout", "origin/master")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
