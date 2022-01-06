package common

import (
	. "github.com/klauspost/cpuid/v2"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Perm struct {
	Digests []string `json:"digests"`
	PermOptions
}

type PermOptions struct {
	Log         *logrus.Logger
	CheckSource string
}

func InitPerm(opt PermOptions) *Perm {
	var perm Perm
	perm.PermOptions = opt
	if perm.Log == nil {
		perm.Log = logrus.New()
	}
	if opt.CheckSource == "" {
		perm.Log.Warn("perm check source not set")
	}
	return &perm
}

func (p *Perm) Confirm() {
	id := GetID()
	if p.Digests == nil {
		p.Log.Error("[SYS] permissions list not found, plz check your network")
		exitOops()
	}
	for _, v := range p.Digests {
		if id == v {
			p.Log.Infof("[SYS] permissions key -> %s [verified]", id)
			return
		}
	}
	p.Log.Warn("[SYS] permissions key -> %s [unverified]", id)
	p.Log.Warn("[SYS] oops, you don't have permission. plz contact the developer (⑉･̆-･̆⑉)")
	exitOops()
}

func exitOops() {
	time.Sleep(time.Second * 5)
	os.Exit(999)
}

func GetID() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("[SYS] panic" + err.Error())
	}
	add := ""
	for _, inter := range interfaces {
		if inter.Name == "以太网" {
			add += inter.HardwareAddr.String()
		}
	}
	add += mixCPUInfo()
	return GetMD5(add)
}

func mixCPUInfo() string {
	v := strings.Join(CPU.FeatureSet(), ",")
	r := CPU.PhysicalCores + CPU.ThreadsPerCore + CPU.LogicalCores + CPU.Family + CPU.Model + CPU.CacheLine + CPU.Cache.L1D*CPU.Cache.L1I
	t := strconv.Itoa(r<<3 + int(CPU.VendorID))
	return GetMD5(v + t)
}
