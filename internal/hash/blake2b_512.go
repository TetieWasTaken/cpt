package hash

import (
	"hash"

	"golang.org/x/crypto/blake2b"
)

func init() {
	RegisterAlgorithm(externalHashErr{
		name: "blake2b-512",
		new:  func() (hash.Hash, error) { return blake2b.New512(nil) },
	})
}
