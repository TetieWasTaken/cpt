package hash

import (
	"strings"
	"testing"
)

// knownHashes are the hashes of the string "The quick brown fox jumps over the lazy dog"
var knownHashes = map[string]string{"sha256": "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592", "md5": "9e107d9d372bb6826bd81d3542a419d6"}

func TestGetAlgorithm(t *testing.T) {
	_, ok := GetAlgorithm("sha256")

	if !ok {
		t.Errorf("GetAlgorithm() was unable to find sha256")
	}
}

func TestHash(t *testing.T) {
	for algorithm, expectedHash := range knownHashes {
		hasher, ok := GetAlgorithm(algorithm)

		if !ok {
			t.Errorf("GetAlgorithm() was unable to find %s", algorithm)
		}

		result, error := hasher.Hash(strings.NewReader("The quick brown fox jumps over the lazy dog"))

		if error != nil {
			t.Errorf("Hash() resulted in an error")
		}

		hex := HexDigest(result)

		if hex != expectedHash {
			t.Errorf("Hash result did not match known hash")
		}
	}
}
