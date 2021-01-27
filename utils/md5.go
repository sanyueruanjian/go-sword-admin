package utils

import (
	"crypto/md5"
	"encoding/hex"

	"project/utils/config"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	value = value + config.JwtConfig.Secret
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
