package passphrases

import (
	"bytes"
	"fmt"
	"io"
	"syscall"

	"golang.org/x/term"
)

// New requests a new passphrase from the user
func New(w io.Writer) (string, error) {
	if _, err := fmt.Fprintf(w, "Enter secret passphrase: "); err != nil {
		return "", fmt.Errorf("failed to write to output: %w", err)
	}

	passphrase, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read secret passphrase from input: %w", err)
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return "", fmt.Errorf("failed to write to output: %w", err)
	}

	if _, err := fmt.Fprintf(w, "Enter secret passphrase again: "); err != nil {
		return "", fmt.Errorf("failed to write to output: %w", err)
	}

	passphraseConfirm, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read secret passphrase from input: %w", err)
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return "", fmt.Errorf("failed to write to output: %w", err)
	}

	if !bytes.Equal(passphrase, passphraseConfirm) {
		return "", fmt.Errorf("passphrases do not match")
	}

	return string(passphrase), nil
}

// Prompt prompts for passphrase entry
func Prompt(w io.Writer) (string, error) {
	if _, err := fmt.Fprintf(w, "Enter secret passphrase: "); err != nil {
		return "", fmt.Errorf("failed to write to output: %w", err)
	}

	passphrase, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read secret passphrase from input: %w", err)
	}

	return string(passphrase), nil
}
