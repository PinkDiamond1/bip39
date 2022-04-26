package prompts

import (
	"fmt"
	"os"
)

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
