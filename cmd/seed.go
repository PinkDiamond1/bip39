/*
Copyright © 2022 kubetrail.io authors

*/
package cmd

import (
	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/run"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "generate hex seed from mnemonic",
	Long: `This command generates hex seed from mnemonic and optional passphrase.

Hex seed is directly used for key generation and should therefore be treated as
a secret material.

Please note that the seed is always generated for the English equivalent mnemonic
if the input mnemonic language is different from English.

The language needs to be one of the following:
1. English (default)
2. Japanese
3. ChineseSimplified
4. ChineseTraditional
5. Czech
6. French
7. Italian
8. Korean
9. Spanish

BIP-39 proposal: https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

Mnemonics are always reformatted using sentence fields and are not affected
by extra white spaces. They are, however, case sensitive. For instance, following
mnemonics are all the same. Note, how white spaces at the beginning, end or
in between fields are ignored.

"farm employ cup erosion half birth become love excite private swallow grit",
"farm employ cup erosion half    birth become love excite private swallow grit",
"farm employ cup erosion half birth become love excite private swallow grit    ",
"    farm employ cup erosion half birth become love excite private swallow grit",
"    farm employ cup     erosion half birth become love   excite private swallow grit    ",
`,
	RunE: run.Seed,
}

func init() {
	rootCmd.AddCommand(seedCmd)
	f := seedCmd.Flags()
	f.String(flags.Language, "English", "Language")
	f.Bool(flags.UsePassphrase, false, "Use passphrase")
	f.Bool(flags.SkipMnemonicValidation, false, "Skip mnemonic validation")
	f.Bool(flags.Short, false, "Generate 32 bytes seed instead of 64")
}
