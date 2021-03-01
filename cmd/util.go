package cmd

import (
	"os"
	"os/exec"
)

func git(args ...string) {
	runOrPanic(exec.Command("git", args...))
}

func showGit(args ...string) {
	c := exec.Command("git", args...)
	c.Stdout = os.Stdout
	runOrPanic(c)
}

func runOrPanic(c *exec.Cmd) {
	if err := c.Run(); err != nil {
		panic(err)
	}
}
