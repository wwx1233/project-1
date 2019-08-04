package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GetBlockHash(prehash, root []byte, tmp int64, height int) string {
	var buf bytes.Buffer
	buf.Write(prehash)
	buf.Write(root)
	buf.Write(Int64ToByte(tmp))
	buf.Write(IntToByte(height))
	return Hash(buf.Bytes())
}

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
