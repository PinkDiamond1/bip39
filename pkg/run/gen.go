package run

import (
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func Gen(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Length, cmd.Flags().Lookup(flags.Length))
	_ = viper.BindPFlag(flags.Language, cmd.Flags().Lookup(flags.Language))

	length := viper.GetInt(flags.Length)
	language := viper.GetInt(flags.Language)

	switch length {
	case 12, 15, 18, 21, 24:
	default:
		return fmt.Errorf("invalid length, can only be 12, 15, 18, 21 or 24")
	}
	bitSize := length / 3 * 32

	switch language {
	case 1:
		bip39.SetWordList(wordlists.English)
	case 2:
		bip39.SetWordList(wordlists.Japanese)
	case 3:
		bip39.SetWordList(wordlists.ChineseSimplified)
	case 4:
		bip39.SetWordList(wordlists.ChineseTraditional)
	case 5:
		bip39.SetWordList(wordlists.Czech)
	case 6:
		bip39.SetWordList(wordlists.French)
	case 7:
		bip39.SetWordList(wordlists.Italian)
	case 8:
		bip39.SetWordList(wordlists.Korean)
	case 9:
		bip39.SetWordList(wordlists.Spanish)
	default:
		return fmt.Errorf("invalid language")
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
