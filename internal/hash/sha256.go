package hash

import (
	"crypto"
	_ "crypto/sha256"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha256", hasher: crypto.SHA256})
}
