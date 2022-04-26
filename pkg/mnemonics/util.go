package mnemonics

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func SetLanguage(language string) error {
	switch strings.ToLower(language) {
	case strings.ToLower(LanguageEnglish):
		bip39.SetWordList(wordlists.English)
	case strings.ToLower(LanguageJapanese):
		bip39.SetWordList(wordlists.Japanese)
	case strings.ToLower(LanguageChineseSimplified):
		bip39.SetWordList(wordlists.ChineseSimplified)
	case strings.ToLower(LanguageChineseTraditional):
		bip39.SetWordList(wordlists.ChineseTraditional)
	case strings.ToLower(LanguageCzech):
		bip39.SetWordList(wordlists.Czech)
	case strings.ToLower(LanguageFrench):
		bip39.SetWordList(wordlists.French)
	case strings.ToLower(LanguageItalian):
		bip39.SetWordList(wordlists.Italian)
	case strings.ToLower(LanguageKorean):
		bip39.SetWordList(wordlists.Korean)
	case strings.ToLower(LanguageSpanish):
		bip39.SetWordList(wordlists.Spanish)
	default:
		return fmt.Errorf("invalid mnemonic language: %s, valid languages are %v",
			language,
			validLanguages,
		)
	}

	return nil
}

func New(length int, language string) (string, error) {
	if err := SetLanguage(language); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	switch length {
	case 12, 15, 18, 21, 24:
	default:
		return "", fmt.Errorf("invalid length, can only be 12, 15, 18, 21 or 24")
	}
	bitSize := length / 3 * 32

	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy: %w", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	return mnemonic, nil
}

func FromReader(r io.Reader) (string, error) {
	inputReader := bufio.NewReader(r)
	mnemonic, err := inputReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read from input: %w", err)
	}

	mnemonic = strings.Join(strings.Fields(mnemonic), " ")

	return mnemonic, nil
}

func Translate(mnemonic, fromLanguage, toLanguage string) (string, error) {
	wordList := bip39.GetWordList()
	defer bip39.SetWordList(wordList)

	if err := SetLanguage(fromLanguage); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return "", fmt.Errorf("mnemonic validation failed")
	}

	entropy, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy from mnemonic: %w", err)
	}

	if err := SetLanguage(toLanguage); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	return mnemonic, nil
}

func Seed(mnemonic, passphrase string) []byte {
	return bip39.NewSeed(mnemonic, passphrase)
}

func Prompt(w io.Writer) error {
	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	if !prompt {
		return nil
	}

	if _, err := fmt.Fprintf(w, "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}

func Validate(mnemonic, language string) error {
	if err := SetLanguage(language); err != nil {
		return fmt.Errorf("failed to set language: %w", err)
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("failed to validate mnemonic for language %s", language)
	}

	return nil
}
