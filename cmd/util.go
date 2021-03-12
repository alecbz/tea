package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var (
	start           time.Time
	delaySpinnerFor = 5 * time.Second
)

func init() {
	start = time.Now()
}

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

func firstLine(message string) string {
	nl := strings.Index(message, "\n")
	if nl != -1 {
		return message[:nl]
	}
	return message
}

func sameWeek(a, b time.Time) bool {
	y1, w1 := a.ISOWeek()
	y2, w2 := b.ISOWeek()
	return y1 == y2 && w1 == w2
}

func humanTime(t time.Time) string {
	t = t.In(time.Local)

	now := time.Now()
	sameYear := t.Year() == now.Year()
	since := time.Since(t)

	// Today
	if sameYear && t.YearDay() == now.YearDay() {
		return humanDuration(since)
	}

	// Yesterday
	if sameYear && t.YearDay() == now.YearDay()-1 {
		return t.Format("yesterday @ 3:04 PM")
	}

	if sameWeek(now, t) {
		return fmt.Sprintf("%s (%s)", t.Format("Monday @ 3:04 PM"), humanDuration(since))
	}

	// Within a ~week
	if since < 7*24*time.Hour {
		return fmt.Sprintf("%s (%s)", t.Format("January 2 @ 3:04 PM"), humanDuration(since))
	}

	// Within a ~month
	if since < 30*24*time.Hour {
		return fmt.Sprintf("%s (%s)", t.Format("January 2 @ 3:04 PM"), humanDuration(since))
	}

	if sameYear {
		return t.Format("January 2")
	}

	return t.Format("January 2, 2006")

	//	Mon Jan 2 15:04:05 -0700 MST 2006
}

func humanDuration(d time.Duration) string {
	if d < time.Minute {
		return "just now"
	}
	if d < time.Hour {
		return fmt.Sprintf("%s ago", pluralize(d/time.Minute, "minutes"))
	}
	if d < 48*time.Hour {
		return fmt.Sprintf("%s ago", pluralize(d/time.Hour, "hours"))
	}
	return fmt.Sprintf("%s ago", pluralize(d/(24*time.Hour), "days"))
}
