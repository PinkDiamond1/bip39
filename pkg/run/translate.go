package run

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func Translate(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.FromLanguage, cmd.Flags().Lookup(flags.FromLanguage))
	_ = viper.BindPFlag(flags.ToLanguage, cmd.Flags().Lookup(flags.ToLanguage))

	fromLanguage := viper.GetInt(flags.FromLanguage)
	toLanguage := viper.GetInt(flags.ToLanguage)

	switch fromLanguage {
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
		return fmt.Errorf("invalid from language")
	}

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	inputReader := bufio.NewReader(cmd.InOrStdin())
	mnemonic, err := inputReader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read from input: %w", err)
	}
	mnemonic = strings.Trim(mnemonic, "\n")

	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("invalid mnemonic")
	}

	entropy, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return fmt.Errorf("failed to generate entropy from mnemonic: %w", err)
	}

	switch toLanguage {
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
		return fmt.Errorf("invalid to language")
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
