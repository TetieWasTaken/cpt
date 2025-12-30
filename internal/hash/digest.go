package hash

import (
	"encoding/base64"
	"encoding/hex"
)

func HexDigest(sum []byte) string {
	return hex.EncodeToString(sum)
}

func Base64Digest(sum []byte) string {
	return base64.StdEncoding.EncodeToString(sum)
}
