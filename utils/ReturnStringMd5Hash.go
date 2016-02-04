package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ReturnStringMd5Hash(input string) string {

	result := md5.Sum([]byte(input))
	return string(hex.EncodeToString(result[:]))

}
