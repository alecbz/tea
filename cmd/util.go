package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

var (
	start           time.Time
	delaySpinnerFor = 5 * time.Second
)

func init() {
	start = time.Now()
}

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

func fetch(remote, branch string) {
	spin(fmt.Sprintf("Fetching %s from %s", branch, remote), func() {
		git("fetch", remote, branch)
	})
}

func spin(msg string, work func()) {
	done := make(chan (struct{}))
	go func() {
		work()
		close(done)
	}()

	var s *spinner.Spinner
	delay := time.After(delaySpinnerFor - time.Since(start))

	for {
		select {
		case <-done:
			if s != nil {
				s.Stop()
			}
			return
		case <-delay:
			s = spinner.New(spinner.CharSets[14], 50*time.Millisecond)
			s.Suffix = " " + msg
			s.Start()
		}
	}
}
