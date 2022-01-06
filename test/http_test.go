package test

import (
	"github.com/r3inbowari/common"
	"testing"
)

func TestHttp(t *testing.T) {
	var res common.TaobaoBody
	common.RequestJson(common.RequestOptions{Url: "https://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"}, &res)
	println(res.Data.T)
}
