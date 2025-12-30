package hash

import "golang.org/x/crypto/sha3"

func init() {
	RegisterAlgorithm(externalHash{name: "sha3-384", new: sha3.New384})
}
