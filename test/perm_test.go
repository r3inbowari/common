package test

import (
	"github.com/r3inbowari/common"
	"testing"
)

func TestPerm(t *testing.T) {
	p := common.InitPermClient(common.PermOptions{CheckSource: "http://127.0.0.1:18833", AppId: "acd3f8c51a"})
	p.Confirm()
}
