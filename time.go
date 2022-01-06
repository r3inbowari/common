package common

import (
	"math/rand"
	"time"
)

type TaobaoBody struct {
	API  string   `json:"api"`
	V    string   `json:"v"`
	Ret  []string `json:"ret"`
	Data Data     `json:"data"`
}

type Data struct {
	T string `json:"t"`
}

func GetTime() *TaobaoBody {
	var ret TaobaoBody
	RequestJson(RequestOptions{Url: "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"}, &ret)
	return &ret
}

func Random(max int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(max)) * time.Second
}
