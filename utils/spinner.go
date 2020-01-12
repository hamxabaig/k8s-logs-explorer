package utils

import (
	"time"

	"github.com/janeczku/go-spinner"
)

func NewSpinner(text string) *spinner.Spinner {
	s := spinner.NewSpinner(text)
	s.SetSpeed(100 * time.Millisecond)
	s.SetCharset([]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"})
	return s
}
