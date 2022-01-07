package main

import (
	"github.com/r3inbowari/common"
	"time"
)

func main() {
	p := common.InitPermClient(common.PermOptions{CheckSource: "https://1077739472743245.cn-hangzhou.fc.aliyuncs.com/2016-08-15/proxy/perm.LATEST/perm", AppId: "acd3f8c51a", ExpireAfter: time.Minute * 2})
	//p := common.InitPermClient(common.PermOptions{CheckSource: "http://127.0.0.1:18833", AppId: "acd3f8c51a", ExpireAfter: time.Hour * 1})
	p.Verify()
}
