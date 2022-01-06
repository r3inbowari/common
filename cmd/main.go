package main

import (
	"context"
	"fmt"
	"github.com/Dreamacro/clash/adapter/outbound"
	"github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/listener/socks"
	"github.com/r3inbowari/common"
	"io"
	"net"
	"time"
)

func main() {

	in := make(chan constant.ConnContext, 100)
	defer close(in)

	//l, err := http.New("127.0.0.1:6666", in)
	l, err := socks.New("127.0.0.1:6666", in)

	if err != nil {
		panic(err)
	}
	defer l.Close()

	println("listen at:", l.Address())

	ss, err := outbound.NewShadowSocks(outbound.ShadowSocksOption{
		BasicOption: outbound.BasicOption{},
		Name:        "home",
		Server:      "invenleey.oicp.net",
		Port:        26185,
		Password:    "159463",
		Cipher:      "chacha20-ietf-poly1305",
		UDP:         false,
	})

	go func() {
		time.Sleep(time.Second * 2)

		common.InitDB()
	}()

	for c := range in {
		conn := c
		metadata := conn.Metadata()
		fmt.Printf("request incoming from %s to %s\n", metadata.SourceAddress(), metadata.RemoteAddress())
		go func() {
			remote, err := ss.DialContext(context.Background(), metadata)
			if err != nil {
				fmt.Printf("dial error: %s\n", err.Error())
				return
			}
			relay(remote, conn.Conn())
		}()
	}
}

func relay(l, r net.Conn) {
	go io.Copy(l, r)
	io.Copy(r, l)
}
