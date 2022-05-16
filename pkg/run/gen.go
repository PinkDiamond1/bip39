package run

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip39"
	"gopkg.in/yaml.v3"
)

type output struct {
	Mnemonic string `json:"mnemonic,omitempty" yaml:"mnemonic,omitempty"`
	Entropy  string `json:"entropy,omitempty" yaml:"entropy,omitempty"`
	Seed     string `json:"seed,omitempty" yaml:"seed,omitempty"`
}

func Gen(cmd *cobra.Command, args []string) error {
	persistentFlags := getPersistentFlags(cmd)

	_ = viper.BindPFlag(flags.Length, cmd.Flag(flags.Length))
	_ = viper.BindPFlag(flags.Language, cmd.Flag(flags.Language))
	_ = viper.BindPFlag(flags.Entropy, cmd.Flag(flags.Entropy))

	length := viper.GetInt(flags.Length)
	language := viper.GetString(flags.Language)
	entropy := viper.GetString(flags.Entropy)

	var mnemonic string
	var err error

	if len(entropy) == 0 && len(args) > 0 {
		entropy = args[0]
	}

	if len(entropy) > 0 {
		b, err := hex.DecodeString(entropy)
		if err != nil {
			return fmt.Errorf("failed to decode entropy: %w", err)
		}

		mnemonic, err = mnemonics.NewFromEntropy(b, language)
		if err != nil {
			return fmt.Errorf("failed to generate new mnemonic from entropy: %w", err)
		}
	} else {
		mnemonic, err = mnemonics.New(length, language)
		if err != nil {
			return fmt.Errorf("failed to generate a new mnemonic: %w", err)
		}
	}

	b, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return fmt.Errorf("failed to generate entropy from mnemonic: %w", err)
	}

	entropy = hex.EncodeToString(b)

	out := &output{
		Mnemonic: mnemonic,
		Entropy:  entropy,
	}

	switch strings.ToLower(persistentFlags.OutputFormat) {
	case flags.OutputFormatNative:
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), mnemonic); err != nil {
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
