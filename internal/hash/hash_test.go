package hash

import (
	"strings"
	"testing"
)

func TestGetAlgorithm(t *testing.T) {
	_, ok := GetAlgorithm("sha256")

	if !ok {
		t.Errorf("GetAlgorithm() was unable to find sha256")
	}
}

func TestHash(t *testing.T) {
	hasher, ok := GetAlgorithm("sha256")

	if !ok {
		t.Errorf("GetAlgorithm() was unable to find sha256")
	}

	result, error := hasher.Hash(strings.NewReader("The quick brown fox jumps over the lazy dog"))

	if error != nil {
		t.Errorf("Hash() resulted in an error")
	}

	hex := HexDigest(result)

	if hex != "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592" {
		t.Errorf("Hash result did not match known hash")
	}
}
