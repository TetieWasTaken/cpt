// Package hash implements all the hashing algorithms and their helper utilities.
package hash

import (
	"crypto"
	"fmt"
	"hash"
	"io"
	"sort"
)

type HashingAlgorithm interface {
	Name() string
	Hash(data io.Reader) ([]byte, error)
}

var algorithms = map[string]HashingAlgorithm{}

func RegisterAlgorithm(hasher HashingAlgorithm) {
	if hasher == nil {
		panic("Unable to register nil hasher.")
	}

	name := hasher.Name()

	if name == "" {
		panic("Unable to register hasher without name")
	}

	if _, exists := algorithms[name]; exists {
		panic(fmt.Sprintf("Unable to register duplicate hasher (%s)", name))
	}

	algorithms[name] = hasher
}

func GetAlgorithm(name string) (HashingAlgorithm, bool) {
	hasher, ok := algorithms[name]

	return hasher, ok
}

func ListAlgorithms() []string {
	res := make([]string, 0, len(algorithms))

	for k := range algorithms {
		res = append(res, k)
	}

	sort.Strings(res)
	return res
}

type cryptoHash struct {
	name   string
	hasher crypto.Hash
}

func (cryptoHasher cryptoHash) Name() string { return cryptoHasher.name }

func (cryptoHasher cryptoHash) Hash(data io.Reader) ([]byte, error) {
	if !cryptoHasher.hasher.Available() {
		return nil, fmt.Errorf("%s is not available", cryptoHasher.name)
	}

	hasher := cryptoHasher.hasher.New()

	if _, error := io.Copy(hasher, data); error != nil {
		return nil, error
	}

	return hasher.Sum(nil), nil
}

type externalHash struct {
	name string
	new  func() hash.Hash
}

func (externalHasher externalHash) Name() string { return externalHasher.name }

func (externalHasher externalHash) Hash(data io.Reader) ([]byte, error) {
	hasher := externalHasher.new()
	if _, err := io.Copy(hasher, data); err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}

type externalHashErr struct {
	name string
	new  func() (hash.Hash, error)
}

func (externalHasher externalHashErr) Name() string { return externalHasher.name }

func (externalHasher externalHashErr) Hash(data io.Reader) ([]byte, error) {
	hasher, err := externalHasher.new()
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(hasher, data); err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
