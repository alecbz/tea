package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func runGit(args ...string) {
	c := exec.Command("git", args...)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("error running %v: %v:\n%s", c, err, string(out))
	}
}

func showGit(args ...string) {
	c := exec.Command("git", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		log.Fatalf("error running %v: %v", c, err)
	}
}

func fetch(remote, branch string) {
	spin(fmt.Sprintf("Fetching %s from %s", branch, remote), func() {
		runGit("fetch", remote, branch)
	})
}

func reflog(repo *git.Repository, ref string) []*object.Commit {
	c := exec.Command("git", "rev-list", "--walk-reflogs", ref)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("error running %v: %v:\n%s", c, err, string(out))
	}

	var commits []*object.Commit

	scan := bufio.NewScanner(bytes.NewReader(out))
	for scan.Scan() {
		c, err := repo.CommitObject(plumbing.NewHash(scan.Text()))
		if err != nil {
			log.Fatalf("error getting commit %q: %v", scan.Text(), err)
		}
		commits = append(commits, c)
	}
	return commits
}
