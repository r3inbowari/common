package common

import (
	"context"
	"fmt"
	proxy_mysql "github.com/go-sql-driver/mysql"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
)

type DBProxyOptions struct {
	User      string `yaml:"user"`
	Pwd       string `yaml:"pwd"`
	Url       string `yaml:"url"`
	Port      string `yaml:"port"`
	Schema    string `yaml:"schema"`
	Socks     string `yaml:"socks"`
	SocksAuth *proxy.Auth
}

func RegisterDBProxy(opts DBProxyOptions) string {
	dialer, err := proxy.SOCKS5("tcp", opts.Socks, nil, proxy.Direct)
	if err != nil {
		return ""
	}

	id := CreateUUID()
	proxy_mysql.RegisterDialContext(id, func(ctx context.Context, addr string) (net.Conn, error) {
		return dialer.Dial("tcp", addr)
	})

	dbUrl := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		opts.User,
		opts.Pwd,
		id,
		opts.Url,
		opts.Port,
		opts.Schema,
	)
	return dbUrl
}

type HttpProxyOptions struct {
	Socks     string `yaml:"socks"`
	SocksAuth *proxy.Auth
}

func RegisterHttpProxy(opts HttpProxyOptions) *http.Client {
	dialer, err := proxy.SOCKS5("tcp", opts.Socks, opts.SocksAuth, proxy.Direct)
	if err != nil {
		return nil
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	httpTransport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial("tcp", addr)
	}
	return httpClient
}
