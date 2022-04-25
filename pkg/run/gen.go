package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
)

func Gen(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Length, cmd.Flag(flags.Length))
	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))

	length := viper.GetInt(flags.Length)
	language := viper.GetString(flags.Language)

	switch length {
	case 12, 15, 18, 21, 24:
	default:
		return fmt.Errorf("invalid length, can only be 12, 15, 18, 21 or 24")
	}
	bitSize := length / 3 * 32

	if err := SetMnemonicLanguage(language); err != nil {
		return fmt.Errorf("failed to set language: %w", err)
	}

	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return fmt.Errorf("failed to generate entropy: %w", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
