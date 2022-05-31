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
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/run"
	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate mnemonic between languages",
	Long: `This command translates mnemonic from one language
to another in the sense that it preserves the underlying entropy:

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
	RunE: run.Translate,
}

func init() {
	rootCmd.AddCommand(translateCmd)
	f := translateCmd.Flags()

	f.String(flags.FromLanguage, mnemonics.LanguageEnglish, "From language")
	f.String(flags.ToLanguage, mnemonics.LanguageEnglish, "To language")

	_ = translateCmd.RegisterFlagCompletionFunc(
		flags.FromLanguage,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					mnemonics.LanguageEnglish,
					mnemonics.LanguageJapanese,
					mnemonics.LanguageChineseSimplified,
					mnemonics.LanguageChineseTraditional,
					mnemonics.LanguageCzech,
					mnemonics.LanguageFrench,
					mnemonics.LanguageItalian,
					mnemonics.LanguageKorean,
					mnemonics.LanguageSpanish,
				},
				cobra.ShellCompDirectiveDefault
		},
	)

	_ = translateCmd.RegisterFlagCompletionFunc(
		flags.ToLanguage,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					mnemonics.LanguageEnglish,
					mnemonics.LanguageJapanese,
					mnemonics.LanguageChineseSimplified,
					mnemonics.LanguageChineseTraditional,
					mnemonics.LanguageCzech,
					mnemonics.LanguageFrench,
					mnemonics.LanguageItalian,
					mnemonics.LanguageKorean,
					mnemonics.LanguageSpanish,
				},
				cobra.ShellCompDirectiveDefault
		},
	)
}
