package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/klauspost/cpuid/v2"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Perm struct {
	PermOptions
	ID       string
	IsExpire bool
}

type PermOptions struct {
	Log         *logrus.Logger
	CheckSource string
	AppId       string
	ExpireAfter time.Duration
}

func InitPermClient(opt PermOptions) *Perm {
	var perm Perm
	perm.PermOptions = opt
	if perm.Log == nil {
		perm.Log = logrus.New()
	}
	if opt.CheckSource == "" {
		perm.Log.Warn("[PERM] check source not set")
	}
	if opt.AppId == "" {
		perm.AppId = "acd3f8c51a"
		perm.Log.WithField("default", perm.AppId).Warn("[PERM] appid not set")

	}
	perm.Log.Info("[PERM] plugins loaded")
	perm.ID = GetID(perm.AppId)
	return &perm
}

func (p *Perm) Verify() {
	ok, err := p.TransportPerm()
	if err != nil {
		p.Log.WithField("err", err.Error()).Error("[PERM] some error happened, plz report it to the developer, thanks!")
	}

	if p.IsExpire {
		p.Log.Warn("[PERM] authorized signature has expired")
	}

	if ok && !p.IsExpire {
		p.Log.Infof("[PERM] permissions key -> %s [verified]", p.ID)
		return
	}

	p.Log.Warnf("[PERM] permissions key -> %s [unverified]", p.ID)
	p.Log.Warn("[PERM] oops, you don't have permission. plz contact the developer (⑉･̆-･̆⑉)")
	exitOops()
}

func exitOops() {
	time.Sleep(time.Second * 5)
	os.Exit(999)
}

func GetID(appid string) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("[PERM] not supported current operating system")
	}
	add := ""
	for _, inter := range interfaces {
		if inter.Name == "以太网" {
			add += inter.HardwareAddr.String()
		}
	}
	add += mixCPUInfo()
	add += appid
	return GetMD5(add)
}

func mixCPUInfo() string {
	v := strings.Join(CPU.FeatureSet(), ",")
	r := CPU.PhysicalCores + CPU.ThreadsPerCore + CPU.LogicalCores + CPU.Family + CPU.Model + CPU.CacheLine + CPU.Cache.L1D*CPU.Cache.L1I
	t := strconv.Itoa(r<<3 + int(CPU.VendorID))
	return GetMD5(v + t)
}

type RequestResult struct {
	Total   int          `json:"total"`
	Data    RSATransport `json:"data"`
	Code    int          `json:"code"`
	Message string       `json:"msg"`
}

func (p *Perm) TransportPerm() (bool, error) {
	url := fmt.Sprintf("%s/verify/%s", p.CheckSource, p.ID)

	pair, err := CreatePair()
	if err != nil {
		return false, errors.New("create pair failed")
	}

	rt := *pair.NewRSATransport()
	rtm, err := json.Marshal(rt)
	if err != nil {
		return false, errors.New("create transport failed")
	}

	var rtr RequestResult

	_, err = RequestJson(RequestOptions{
		Url:    url,
		Reader: bytes.NewBuffer(rtm),
	}, &rtr)
	if err != nil {
		return false, err
	}

	d, err := rtr.Data.Decrypt(pair.Private)
	if err != nil {
		return false, errors.New("decrypt failed")
	}

	if string(d) == "!*#r3inbowari@gmail.com#*b41b4af78240ae375a4cb0b95932ffc3" {
		return false, nil
	}

	var sign Sign
	err = json.Unmarshal(d, &sign)
	if err != nil {
		return false, err
	}

	networkTime := GetTime()
	it, err := strconv.ParseInt(networkTime.Data.T, 10, 64)
	if err != nil {
		return false, err
	}

	fit := sign.Time.Add(p.ExpireAfter).UnixMilli()
	if sign.Time.Location().String() == "UTC" {
		// aliyun fc 默认使用utc时间
		fit -= 28800 * 1000
	}
	p.IsExpire = fit < it
	return fit > it, nil
}

type Sign struct {
	Time   time.Time // 生成时间
	Random string    // 随机数
}
