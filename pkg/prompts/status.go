package prompts

import (
	"fmt"
	"os"
)

// Status checks is STDIN is connected.
// This can be used to switch off prompting
// for user input assuming input is being fed
// via STDIN pipe
func Status() (bool, error) {
	prompt := false
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("failed ot stat stdin")
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		prompt = true
	}

	return prompt, nil
}
