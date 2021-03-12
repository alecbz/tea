package cmd

import (
	"fmt"
	"time"
)

func pluralize(n time.Duration, obj string) string {
	if n == 1 && obj[len(obj)-1] == 's' {
		return fmt.Sprintf("%d %s", n, obj[:len(obj)-1])
	}
	return fmt.Sprintf("%d %s", n, obj)
}
