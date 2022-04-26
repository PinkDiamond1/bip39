/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/kubetrail/bip39/pkg/flags"
	"github.com/kubetrail/bip39/pkg/run"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check validity of mnemonic",
	Long: `Mnemonic is not just any arbitrary list of words.
It has a structure (last word contains checksum) and comes from
a predefined wordlist.

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
"    farm employ cup     erosion half birth become love   excite private swallow grit    ",`,
	RunE: run.Validate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
	f := validateCmd.Flags()
	f.String(flags.Language, "English", "Language")
}
