/*
Copyright © 2022 kubetrail.io authors

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

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new BIP39 mnemonic",
	Long: `This command generates a new BIP39 mnemonic.
The length can be 12, 15, 18, 21 or 24 words

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
`,
	RunE: run.Gen,
	Args: cobra.ExactArgs(0),
}

func init() {
	rootCmd.AddCommand(genCmd)
	f := genCmd.Flags()

	f.Int(flags.Length, 24, "Number of words")
	f.String(flags.Language, "English", "Language")
	f.String(flags.Entropy, "", "Entropy as hex string (overrides length arg)")
}
