package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

// CreateSpinner creates and configures a new spinner instance with the specified message.
// It returns a pointer to the configured spinner.Spinner object.
func CreateSpinner(message string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Color("cyan")
	return s
}
