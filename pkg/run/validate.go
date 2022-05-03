package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Validate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Language, cmd.Flags().Lookup(flags.Language))
	language := viper.GetString(flags.Language)

	var mnemonic string
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

	if err := mnemonics.Validate(mnemonic, language); err != nil {
		return fmt.Errorf("failed to validate mnemonic: %w", err)
	}

	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	if prompt {
		if _, err := fmt.Fprintf(cmd.OutOrStdout(), "mnemonic is valid in %s language\n", language); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	return nil
}
