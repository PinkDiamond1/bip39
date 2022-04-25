package run

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"syscall"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/term"
)

func Seed(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))
	_ = viper.BindPFlag(flags.UsePassphrase, cmd.Flag(flags.UsePassphrase))
	_ = viper.BindPFlag(flags.SkipMnemonicValidation, cmd.Flag(flags.SkipMnemonicValidation))

	language := viper.GetString(flags.Language)
	usePassphrase := viper.GetBool(flags.UsePassphrase)
	skipMnemonicValidation := viper.GetBool(flags.SkipMnemonicValidation)

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	mnemonic, err := MnemonicFromReader(cmd.InOrStdin())
	if err != nil {
		return fmt.Errorf("failed to read mnemonic from input: %w", err)
	}

	if !skipMnemonicValidation {
		mnemonic, err = TranslateMnemonic(mnemonic, language, LanguageEnglish)
		if err != nil {
			return fmt.Errorf("failed to translate mnemonic to English: %w", err)
		}
	}

	var passphrase []byte
	if usePassphrase {
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter secret passphrase: "); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}

		passphrase, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read secret passphrase from input: %w", err)
		}
		if _, err := fmt.Fprintln(cmd.OutOrStdout()); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}

		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter secret passphrase again: "); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}

		passphraseConfirm, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read secret passphrase from input: %w", err)
		}

		if _, err := fmt.Fprintln(cmd.OutOrStdout()); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}

		if !bytes.Equal(passphrase, passphraseConfirm) {
			return fmt.Errorf("passphrases do not match")
		}
	}

	seed := bip39.NewSeed(mnemonic, string(passphrase))
	if _, err := fmt.Fprintln(cmd.OutOrStdout(), hex.EncodeToString(seed)); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
