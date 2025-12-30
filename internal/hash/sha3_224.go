package hash

import "golang.org/x/crypto/sha3"

func init() {
	RegisterAlgorithm(externalHash{name: "sha3-224", new: sha3.New224})
}
