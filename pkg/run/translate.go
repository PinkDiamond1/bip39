package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Translate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.FromLanguage, cmd.Flags().Lookup(flags.FromLanguage))
	_ = viper.BindPFlag(flags.ToLanguage, cmd.Flags().Lookup(flags.ToLanguage))

	fromLanguage := viper.GetString(flags.FromLanguage)
	toLanguage := viper.GetString(flags.ToLanguage)

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	mnemonic, err := MnemonicFromReader(cmd.InOrStdin())
	if err != nil {
		return fmt.Errorf("failed to read mnemonic from input: %w", err)
	}

	mnemonic, err = TranslateMnemonic(mnemonic, fromLanguage, toLanguage)
	if err != nil {
		return fmt.Errorf("failed to translate mnemonic: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
