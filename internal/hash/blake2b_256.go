package hash

import (
	"hash"

	"golang.org/x/crypto/blake2b"
)

func init() {
	RegisterAlgorithm(externalHashErr{
		name: "blake2b-256",
		new:  func() (hash.Hash, error) { return blake2b.New256(nil) },
	})
}
