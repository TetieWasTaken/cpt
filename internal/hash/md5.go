package hash

import (
	"crypto"
	_ "crypto/md5"
)

func init() {
	RegisterAlgorithm(cryptoHash{name: "md5", hasher: crypto.MD5})
}
