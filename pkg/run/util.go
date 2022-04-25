package run

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

func SetMnemonicLanguage(language string) error {
	switch strings.ToLower(language) {
	case LanguageEnglish:
		bip39.SetWordList(wordlists.English)
	case LanguageJapanese:
		bip39.SetWordList(wordlists.Japanese)
	case LanguageChineseSimplified:
		bip39.SetWordList(wordlists.ChineseSimplified)
	case LanguageChineseTraditional:
		bip39.SetWordList(wordlists.ChineseTraditional)
	case LanguageCzech:
		bip39.SetWordList(wordlists.Czech)
	case LanguageFrench:
		bip39.SetWordList(wordlists.French)
	case LanguageItalian:
		bip39.SetWordList(wordlists.Italian)
	case LanguageKorean:
		bip39.SetWordList(wordlists.Korean)
	case LanguageSpanish:
		bip39.SetWordList(wordlists.Spanish)
	default:
		return fmt.Errorf("invalid mnemonic language: %s", language)
	}

	return nil
}

func MnemonicFromReader(r io.Reader) (string, error) {
	inputReader := bufio.NewReader(r)
	mnemonic, err := inputReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read from input: %w", err)
	}

	return strings.Join(strings.Fields(mnemonic), " "), nil
}

func TranslateMnemonic(mnemonic, fromLanguage, toLanguage string) (string, error) {
	wordList := bip39.GetWordList()
	defer bip39.SetWordList(wordList)

	if err := SetMnemonicLanguage(fromLanguage); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return "", fmt.Errorf("mnemonic validation failed")
	}

	entropy, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy from mnemonic: %w", err)
	}

	if err := SetMnemonicLanguage(toLanguage); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	return mnemonic, nil
}
