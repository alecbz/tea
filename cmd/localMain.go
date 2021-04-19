package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// localMainCmd represents the localMain command
var localMainCmd = &cobra.Command{
	Use:   "local-main",
	Short: "Check out a local main branch",
	Run: func(cmd *cobra.Command, args []string) {
		fetch("origin", config.MainBranch)

		r, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		// Apparently go-git is a steaming pile of dog shit and just omits
		// branches from cfg.Branches for fun sometimes.
		cfg, err := r.Config()
		if err != nil {
			log.Fatal("fuck you: ", err)
		}
		fmt.Println("All the fucking branches:", cfg.Branches)

		h, err := r.Head()
		if err != nil {
			log.Fatal(err)
		}
		currentBranch := strings.TrimPrefix(string(h.Name()), "refs/heads/")

		if currentBranch == config.MainBranch {
			runGit("checkout", fmt.Sprintf("origin/%s", config.MainBranch))
		}

		// re-open after a manual operation, maybe?
		r, err = git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		err = r.DeleteBranch(config.MainBranch)
		fmt.Println("Err from delete branch:", err)
		if err != nil && err != git.ErrBranchNotFound {
			log.Fatal(err)
		}

		runGit("switch", "--create", config.MainBranch, fmt.Sprintf("origin/%s", config.MainBranch))
	},
}

func init() {
	rootCmd.AddCommand(localMainCmd)
}
