package test

import (
	"fmt"
	"github.com/r3inbowari/common"
	"testing"
)

func TestRsa(t *testing.T) {

	pair, err := common.CreatePair()
	if err != nil {
		return
	}

	ppt := pair.NewRSATransport()
	_ = ppt.Encrypt([]byte("Hello"))
	a, _ := ppt.Decrypt(pair.Private)
	fmt.Printf("Encrypt: %b\n", ppt.M)
	fmt.Printf("Decrypt: %s\n", a)
}
