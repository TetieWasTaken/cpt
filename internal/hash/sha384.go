package hash

import (
	"crypto"
	_ "crypto/sha512"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "sha384", hasher: crypto.SHA384})
}
