package test

import (
	"github.com/r3inbowari/common"
	"testing"
)

func TestPerm(t *testing.T) {
	p := common.InitPerm(common.PermOptions{})
	p.Confirm()

}
