package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
)

func Validate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Language, cmd.Flags().Lookup(flags.Language))
	language := viper.GetString(flags.Language)

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	mnemonic, err := MnemonicFromReader(cmd.InOrStdin())
	if err != nil {
		return fmt.Errorf("failed to read mnemonic from input: %w", err)
	}

	if err := SetMnemonicLanguage(language); err != nil {
		return fmt.Errorf("failed to set language: %w", err)
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("invalid mnemonic")
	}

	return nil
}
