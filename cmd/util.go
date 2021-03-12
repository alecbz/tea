package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
)

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
