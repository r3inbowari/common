package test

import (
	"github.com/r3inbowari/common"
	"testing"
)

func TestPerm(t *testing.T) {
	p := common.InitPermClient(common.PermOptions{CheckSource: "https://1077739472743245.cn-hangzhou.fc.aliyuncs.com/2016-08-15/proxy/perm.LATEST/perm", AppId: "acd3f8c51a"})
	p.Confirm()
}
