package hash

import (
	"crypto"
	_ "crypto/sha1"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha1", hasher: crypto.SHA1})
}
