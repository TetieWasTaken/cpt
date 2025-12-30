package hash

import "golang.org/x/crypto/sha3"

func init() {
	RegisterAlgorithm(externalHash{name: "sha3-256", new: sha3.New256})
}
