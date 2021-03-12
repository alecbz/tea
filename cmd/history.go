package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
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
		r, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		head, err := r.Head()
		if err != nil {
			log.Fatal(err)
		}
		branch := strings.TrimPrefix(string(head.Name()), "refs/heads/")

		p := page()
		for _, c := range reflog(r, branch) {
			if p.Done() {
				break
			}
			fmt.Fprintf(p, "%s %s \n", hash(c.Hash.String()[:7]), message(firstLine(c.Message)))
			fmt.Fprintf(p, "\t%s - %s\n\n", author(person(c.Author)), humanTime(c.Author.When))
		}
		p.Wait()

		// showGit("reflog", strings.TrimSpace(string(branch)), "--date=relative")
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
