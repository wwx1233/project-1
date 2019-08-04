package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func IntToByte(num int) []byte {
	return []byte(strconv.Itoa(num))
}

func Int64ToByte(num int64) []byte {
	return []byte(strconv.FormatInt(num, 10))
}

func Hash(b []byte) string {
	hash := sha256.New()
	hash.Write(b)
	str := hash.Sum(nil)
	return hex.EncodeToString(str)
}
