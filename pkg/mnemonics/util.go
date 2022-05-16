package mnemonics

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
)

// SetLanguage sets language for mnemonic generation and validation
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

// New generates a new mnemonic of specified length and for
// requested input language.
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

// NewFromFields generates a new mnemonic sentence using input
// fields
func NewFromFields(fields []string) string {
	return strings.Join(
		strings.Fields(
			strings.Join(fields, " "),
		),
		" ",
	)
}

// NewFromEntropy generates new mnemonic sequence from a valid
// length of entropy bytes. Valid lengths are
// 16, 20, 24, 28, 32.
func NewFromEntropy(entropy []byte, language string) (string, error) {
	if err := SetLanguage(language); err != nil {
		return "", fmt.Errorf("failed to set language: %w", err)
	}

	switch len(entropy) {
	case 16, 20, 24, 28, 32:
	default:
		return "", fmt.Errorf("entropy length must be 16, 20, 24, 28 or 32, got %d", len(entropy))
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate new mnemonic: %w", err)
	}

	return mnemonic, nil
}

// Prompt prompts for new mnemonic
func Prompt(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "Enter mnemonic: "); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	return nil
}

// Read reads a new mnemonic from input. It splits
// input into fields and then recombines them to ensure
// that the mnemonic is not affected by extra whitespaces
// and/or trailing newline char
func Read(r io.Reader) (string, error) {
	inputReader := bufio.NewReader(r)
	mnemonic, err := inputReader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read from input: %w", err)
	}

	mnemonic = strings.Join(strings.Fields(mnemonic), " ")

	return mnemonic, nil
}

// Translate translates mnemonic in a way that the underlying entropy
// is preserved
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

// Validate validates mnemonic with default language to be english.
// please note that despite language input being variadic, at most
// only one value is accepted
func Validate(mnemonic string, language ...string) error {
	l := LanguageEnglish

	if len(language) > 1 {
		return fmt.Errorf("mnemonic can only be validated for one language, received %d", len(language))
	}

	if len(language) > 0 {
		l = language[0]
	}

	if err := SetLanguage(l); err != nil {
		return fmt.Errorf("failed to set language: %w", err)
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("failed to validate mnemonic for language %s", l)
	}

	return nil
}

// Tidy ensures that the mnemonic sentence does not have extra
// white spaces or trailing newlines. Since mnemonic sentence
// is used in cryptographic seed generation it is necessary
// to make sure that same seed is generated if the fields
// are the same irrespective of spacing between them
func Tidy(mnemonic string) string {
	return strings.Join(strings.Fields(mnemonic), " ")
}
