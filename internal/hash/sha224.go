package hash

import (
	"crypto"
	_ "crypto/sha256"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha224", hasher: crypto.SHA224})
}
