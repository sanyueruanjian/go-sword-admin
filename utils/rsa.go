package utils

import (
	"errors"
	"github.com/wenzhenxi/gorsa"
)

var publicKey =
`-----BEGIN 公钥-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANL378k3RiZHWx5AfJqdH9xRNBmD9wGD2iRe41HdTNF8RUhNnHit5NpMNtGL0NPTSSpPjjI1kJfVorRvaQerUgkCAwEAAQ==
-----END 公钥-----
`

var privateKey =
`-----BEGIN 私钥-----
MIIBUwIBADANBgkqhkiG9w0BAQEFAASCAT0wggE5AgEAAkEA0vfvyTdGJkdbHkB8mp0f3FE0GYP3AYPaJF7jUd1M0XxFSE2ceK3k2kw20YvQ09NJKk+OMjWQl9WitG9pB6tSCQIDAQABAkA2SimBrWC2/wvauBuYqjCFwLvYiRYqZKThUS3MZlebXJiLB+Ue/gUifAAKIg1avttUZsHBHrop4qfJCwAI0+YRAiEA+W3NK/RaXtnRqmoUUkb59zsZUBLpvZgQPfj1MhyHDz0CIQDYhsAhPJ3mgS64NbUZmGWuuNKp5coY2GIj/zYDMJp6vQIgUueLFXv/eZ1ekgz2Oi67MNCk5jeTF2BurZqNLR3MSmUCIFT3Q6uHMtsB9Eha4u7hS31tj1UWE+D+ADzp59MGnoftAiBeHT7gDMuqeJHPL4b+kC+gzV4FGTfhR9q3tTbklZkD2A==
-----END 私钥-----
`

// 私钥解密
func RsaPriDecode(str string) (value string, err error) {
	value, err = gorsa.PriKeyDecrypt(str,privateKey)
	if err != nil {
		return
	}
	return
}

// 公钥解密
func RsaPubDecode(str string) (value string, err error) {
	value, err = gorsa.PublicDecrypt(str, publicKey)
	if err != nil {
		return
	}
	return
}

// 私钥加密
func RsaPriEncode(str string) (value string, err error) {
	value, err = gorsa.PriKeyEncrypt(str, privateKey)
	if err != nil {
		return
	}
	return
}

// 公钥加密
func RsaPubEncode(str string) (value string, err error) {
	value, err = gorsa.PublicEncrypt(str, publicKey)
	if err != nil {
		return
	}
	return
}

// 公钥加密私钥解密
func ApplyPubEPriD() error {
	pubenctypt, err := gorsa.PublicEncrypt(`hello world`,publicKey)
	if err != nil {
		return err
	}

	pridecrypt, err := gorsa.PriKeyDecrypt(pubenctypt,privateKey)
	if err != nil {
		return err
	}
	if string(pridecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

// 公钥解密私钥加密
func ApplyPriEPubD() error {
	prienctypt, err := gorsa.PriKeyEncrypt(`hello world`, privateKey)
	if err != nil {
		return err
	}

	pubdecrypt, err := gorsa.PublicDecrypt(prienctypt, publicKey)
	if err != nil {
		return err
	}
	if string(pubdecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}
