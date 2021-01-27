package utils

import (
	"fmt"
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	md5 := EncodeMD5("123456")
	fmt.Println(md5)
}
