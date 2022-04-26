package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Gen(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Length, cmd.Flag(flags.Length))
	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))

	length := viper.GetInt(flags.Length)
	language := viper.GetString(flags.Language)

	mnemonic, err := mnemonics.New(length, language)
	if err != nil {
		return fmt.Errorf("failed to generate a new mnemonic: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
