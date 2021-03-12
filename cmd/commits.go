package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/spf13/cobra"
)

var (
	hash    = color.New(color.FgYellow).SprintFunc()
	author  = color.New(color.FgBlue).SprintFunc()
	message = color.New(color.FgGreen).SprintFunc()
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
			if p.Done() {
				return storer.ErrStop
			}
			fmt.Fprintf(p, "%s %s \n", hash(c.Hash.String()[:7]), message(firstLine(c.Message)))
			fmt.Fprintf(p, "\t%s - %s\n\n", author(person(c.Author)), humanTime(c.Author.When))
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
