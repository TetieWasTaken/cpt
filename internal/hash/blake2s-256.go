package hash

import (
	"hash"

	"golang.org/x/crypto/blake2s"
)

func init() {
	RegisterAlgorithm(externalHashErr{
		name: "blake2s-256",
		new:  func() (hash.Hash, error) { return blake2s.New256(nil) },
	})
}
