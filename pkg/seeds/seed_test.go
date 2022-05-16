package seeds

import (
	"encoding/hex"
	"testing"
)

const (
	mnemonic   = "ghost alone media humble ride curious among element oval feature donate level"
	passphrase = "pc9S]jO2Q/Fw84R%"
)

func TestNewWithPassphrase(t *testing.T) {
	expectedSeed := "2b6c98c91e473c384310fafbf3ba762ccb196a44eb097da96b0454fa824753b384e35484d10cd634165cf949b1047f6469dbb7fa0366ba43a23d966f36405604"
	if hex.EncodeToString(New(mnemonic, passphrase)) != expectedSeed {
		t.Fatal("seed is not as expected")
	}
}

func TestNewWithoutPassphrase(t *testing.T) {
	expectedSeed := "a60644101e7c9466a14b10cda9d3eece30eb6bc8b1ad06724f6da06d7404b7581d2207c9782b81a058fcac526c5063f9b9bff41d2d2d1e0edcbeced5d097dd45"
	if hex.EncodeToString(New(mnemonic, "")) != expectedSeed {
		t.Fatal("seed is not as expected")
	}
}
