package cmd

import (
	"io"
	"os"
	"os/exec"
)

type pager struct {
	w    *io.PipeWriter
	done chan struct{}
}

func page() pager {
	r, w := io.Pipe()
	done := make(chan struct{})

	less := exec.Command("less")
	less.Stdin = r
	less.Stdout = os.Stdout
	less.Stderr = os.Stderr

	go func() {
		defer close(done)
		defer r.Close()
		less.Run()
	}()

	return pager{w: w, done: done}
}

func (p pager) Write(b []byte) (int, error) {
	return p.w.Write(b)
}

func (p pager) Wait() {
	p.w.Close()
	<-p.done
}
