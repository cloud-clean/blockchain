package common

import "crypto/sha256"

func GetHash(data []byte)[]byte{
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
