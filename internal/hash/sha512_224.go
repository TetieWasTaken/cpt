package hash

import (
	"crypto"
	_ "crypto/sha512"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha512-224", hasher: crypto.SHA512_224})
}
