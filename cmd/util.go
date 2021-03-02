package cmd

import (
	"log"
	"os"
	"os/exec"
)

func git(args ...string) {
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
