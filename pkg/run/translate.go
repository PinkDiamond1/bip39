package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Translate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.FromLanguage, cmd.Flags().Lookup(flags.FromLanguage))
	_ = viper.BindPFlag(flags.ToLanguage, cmd.Flags().Lookup(flags.ToLanguage))

	fromLanguage := viper.GetString(flags.FromLanguage)
	toLanguage := viper.GetString(flags.ToLanguage)

	var mnemonic string
	var err error
	if len(args) == 0 {
		prompt, err := prompts.Status()
		if err != nil {
			return fmt.Errorf("failed to get prompt status: %w", err)
		}

		if prompt {
			if err := mnemonics.Prompt(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to prompt for mnemonic: %w", err)
			}
		}

		mnemonic, err = mnemonics.Read(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read mnemonic from input: %w", err)
		}
	} else {
		mnemonic = mnemonics.NewFromFields(args)
	}

	mnemonic, err = mnemonics.Translate(mnemonic, fromLanguage, toLanguage)
	if err != nil {
		return fmt.Errorf("failed to translate mnemonic: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
