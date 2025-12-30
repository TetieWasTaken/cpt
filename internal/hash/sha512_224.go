package hash

import (
	"crypto"
	_ "crypto/sha512"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha512_224", hasher: crypto.SHA512_224})
}
