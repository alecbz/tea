package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "Show the commit graph",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		head, err := r.Head()
		if err != nil {
			log.Fatal(err)
		}
		head.Type()

		iter, err := r.Log(&git.LogOptions{From: head.Hash()})
		p := page()
		err = iter.ForEach(func(c *object.Commit) error {
			fmt.Fprintf(p, "%s (WTB decorations) \n", yellow(c.Hash.String()[:7]))
			fmt.Fprintf(p, "%s, %s\n", blue(person(c.Author)), humanTime(c.Author.When))
			fmt.Fprintln(p)
			fmt.Fprintf(p, "%s\n", c.Message)
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		p.Wait()
	},
}

func init() {
	rootCmd.AddCommand(commitsCmd)
}
