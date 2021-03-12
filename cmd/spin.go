package cmd

import (
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
