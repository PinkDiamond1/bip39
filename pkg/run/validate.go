package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Validate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Language, cmd.Flags().Lookup(flags.Language))
	language := viper.GetString(flags.Language)

	var mnemonic string
	if len(args) == 0 {
		err := mnemonics.Prompt(cmd.OutOrStdout())
		if err != nil {
			return fmt.Errorf("failed to prompt for mnemonic: %w", err)
		}

		mnemonic, err = mnemonics.FromReader(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read mnemonic from input: %w", err)
		}
	} else {
		mnemonic = mnemonics.FromFields(args)
	}

	if err := mnemonics.Validate(mnemonic, language); err != nil {
		return fmt.Errorf("failed to validate mnemonic: %w", err)
	}

	return nil
}
