package run

import (
	"encoding/hex"
	"fmt"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/passphrases"
	"github.com/kubetrail/bip39/pkg/seeds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Seed(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))
	_ = viper.BindPFlag(flags.UsePassphrase, cmd.Flag(flags.UsePassphrase))
	_ = viper.BindPFlag(flags.SkipMnemonicValidation, cmd.Flag(flags.SkipMnemonicValidation))

	language := viper.GetString(flags.Language)
	usePassphrase := viper.GetBool(flags.UsePassphrase)
	skipMnemonicValidation := viper.GetBool(flags.SkipMnemonicValidation)

	var mnemonic string
	var err error

	if len(args) == 0 {
		err := mnemonics.Prompt(cmd.OutOrStdout())
		if err != nil {
			return fmt.Errorf("failed to prompt for mnemonic: %w", err)
		}

		mnemonic, err = mnemonics.Read(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read mnemonic from input: %w", err)
		}
	} else {
		mnemonic = mnemonics.NewFromFields(args)
	}

	if !skipMnemonicValidation {
		mnemonic, err = mnemonics.Translate(mnemonic, language, mnemonics.LanguageEnglish)
		if err != nil {
			return fmt.Errorf("failed to translate mnemonic to English: %w", err)
		}
	}

	var passPhrase string
	if usePassphrase {
		passPhrase, err = passphrases.New(cmd.OutOrStdout())
		if err != nil {
			return fmt.Errorf("failed to prompt passphrase: %w", err)
		}
	}

	if _, err := fmt.Fprintln(
		cmd.OutOrStdout(),
		hex.EncodeToString(
			seeds.New(mnemonic, passPhrase),
		),
	); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}
