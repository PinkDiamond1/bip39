package run

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/passphrases"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/kubetrail/bip39/pkg/seeds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Seed(cmd *cobra.Command, args []string) error {
	persistentFlags := getPersistentFlags(cmd)

	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))
	_ = viper.BindPFlag(flags.UsePassphrase, cmd.Flag(flags.UsePassphrase))
	_ = viper.BindPFlag(flags.SkipMnemonicValidation, cmd.Flag(flags.SkipMnemonicValidation))
	_ = viper.BindPFlag(flags.Short, cmd.Flag(flags.Short))

	language := viper.GetString(flags.Language)
	usePassphrase := viper.GetBool(flags.UsePassphrase)
	skipMnemonicValidation := viper.GetBool(flags.SkipMnemonicValidation)
	short := viper.GetBool(flags.Short)

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

	seed := seeds.New(mnemonic, passPhrase)
	if len(seed) != seeds.SizeDefault {
		return fmt.Errorf("expected seed len %d, got %d", seeds.SizeDefault, len(seed))
	}

	if short {
		seed = seed[:seeds.SizeShort]
	}

	out := &output{
		Seed: hex.EncodeToString(seed),
	}

	switch strings.ToLower(persistentFlags.OutputFormat) {
	case flags.OutputFormatNative:
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), out.Seed); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	case flags.OutputFormatYaml:
		b, err := yaml.Marshal(out)
		if err != nil {
			return fmt.Errorf("failed to serialize output to yaml: %w", err)
		}

		if _, err := fmt.Fprint(cmd.OutOrStdout(), string(b)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	case flags.OutputFormatJson:
		b, err := json.Marshal(out)
		if err != nil {
			return fmt.Errorf("failed to serialize output to json: %w", err)
		}

		if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(b)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	default:
		return fmt.Errorf("invalid output format")
	}

	return nil
}
