package run

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Validate(cmd *cobra.Command, args []string) error {
	persistentFlags := getPersistentFlags(cmd)

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

	valid := true
	if err := mnemonics.Validate(mnemonic, language); err != nil {
		valid = false
	}

	out := &struct {
		Valid bool `json:"valid" yaml:"valid"`
	}{
		Valid: valid,
	}

	switch strings.ToLower(persistentFlags.OutputFormat) {
	case flags.OutputFormatNative:
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), out.Valid); err != nil {
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
