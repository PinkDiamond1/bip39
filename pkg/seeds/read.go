package seeds

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func Prompt(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "Enter seed in hex: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}

func Read(r io.Reader) ([]byte, error) {
	inputReader := bufio.NewReader(r)
	hexSeed, err := inputReader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read hex seed from input: %w", err)
	}
	hexSeed = strings.Trim(hexSeed, "\n")

	seed, err := hex.DecodeString(hexSeed)
	if err != nil {
		return nil, fmt.Errorf("invalid hex seed string: %w", err)
	}

	return seed, nil
}
