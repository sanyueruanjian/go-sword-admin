package utils

import (
	"fmt"
	"testing"
)

func TestRsaPriDecode(t *testing.T) {
	s, e := RsaPriDecode("n0Woql2nkASUEbcN1wzmx00X/bjIXtBU2nVCHBcJ3usihfS94C7OJ+tcImNNpMjSByrSRz/RA6WRvxpTk2DbxQ==")
	if e != nil {
		fmt.Printf("%v\n", e)
	}
	fmt.Println(s)
}
