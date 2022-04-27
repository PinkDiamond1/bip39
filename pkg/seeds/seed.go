package seeds

import "github.com/tyler-smith/go-bip39"

// New generates a new seed using mnemonic
// and passphrase
func New(mnemonic, passphrase string) []byte {
	return bip39.NewSeed(mnemonic, passphrase)
}
