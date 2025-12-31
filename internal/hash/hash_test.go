package hash

import (
	"strings"
	"testing"
)

// knownHashes are the hashes of the string "The quick brown fox jumps over the lazy dog"
var knownHashes = map[string]string{
	"md5":         "9e107d9d372bb6826bd81d3542a419d6",
	"sha1":        "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12",
	"sha224":      "730e109bd7a8a32b1cb9d9a09aa2325d2430587ddbc0c38bad911525",
	"sha256":      "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592",
	"sha384":      "ca737f1014a48f4c0b6dd43cb177b0afd9e5169367544c494011e3317dbf9a509cb1e5dc1e85a941bbee3d7f2afbc9b1",
	"sha512":      "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6",
	"sha512-224":  "944cd2847fb54558d4775db0485a50003111c8e5daa63fe722c6aa37",
	"sha512-256":  "dd9d67b371519c339ed8dbd25af90e976a1eeefd4ad3d889005e532fc5bef04d",
	"sha3-224":    "d15dadceaa4d5d7bb3b48f446421d542e08ad8887305e28d58335795",
	"sha3-256":    "69070dda01975c8c120c3aada1b282394e7f032fa9cf32f4cb2259a0897dfc04",
	"sha3-384":    "7063465e08a93bce31cd89d2e3ca8f602498696e253592ed26f07bf7e703cf328581e1471a7ba7ab119b1a9ebdf8be41",
	"sha3-512":    "01dedd5de4ef14642445ba5f5b97c15e47b9ad931326e4b0727cd94cefc44fff23f07bf543139939b49128caf436dc1bdee54fcb24023a08d9403f9b4bf0d450",
	"blake2s-256": "606beeec743ccbeff6cbcdf5d5302aa855c256c29b88c8ed331ea1a6bf3c8812",
	"blake2b-256": "01718cec35cd3d796dd00020e0bfecb473ad23457d063b75eff29c0ffa2e58a9",
	"blake2b-384": "b7c81b228b6bd912930e8f0b5387989691c1cee1e65aade4da3b86a3c9f678fc8018f6ed9e2906720c8d2a3aeda9c03d",
	"blake2b-512": "a8add4bdddfd93e4877d2746e62817b116364a1fa7bc148d95090bc7333b3673f82401cf7aa2e4cb1ecd90296e3f14cb5413f8ed77be73045b13914cdcd6a918",
}

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
			t.Fatalf("hash mismatch for %s\n got:  %s\n want: %s", algorithm, hex, expectedHash)
		}
	}
}
