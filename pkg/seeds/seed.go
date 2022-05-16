package seeds

import "github.com/tyler-smith/go-bip39"

const (
	SizeDefault = 64
	SizeShort   = 32
)

// New generates a new seed using mnemonic
// and passphrase
func New(mnemonic, passphrase string) []byte {
	return bip39.NewSeed(mnemonic, passphrase)
}
