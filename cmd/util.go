package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var (
	start           time.Time
	delaySpinnerFor = 5 * time.Second
)

func init() {
	start = time.Now()
}

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
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

func person(s object.Signature) string {
	at := strings.Index(s.Email, "@")
	if at != -1 {
		return s.Email[:at+1]
	}
	return s.Email
}

func humanTime(t time.Time) string {
	t = t.In(time.Local)

	now := time.Now()
	sameYear := t.Year() == now.Year()
	if sameYear && t.Month() == now.Month() && t.Day() == now.Day() {
		return humanDuration(time.Since(t))
	}
	if sameYear {
		//	Mon Jan 2 15:04:05 -0700 MST 2006
		return t.Format("Mon Jan 2, 3:04 PM")
	}
	return t.String()
}

func humanDuration(d time.Duration) string {
	if d < time.Minute {
		return "now"
	}
	if d < time.Hour {
		return fmt.Sprintf("%d minutes ago", d/time.Minute)
	}
	if d < 2*24*time.Hour {
		return fmt.Sprintf("%d hours ago", d/time.Hour)
	}
	return fmt.Sprintf("%d days ago", d/(24*time.Hour))
}
