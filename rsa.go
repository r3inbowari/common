package common

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"math/big"
)

type Pair struct {
	Private *rsa.PrivateKey
}

func CreatePair() (*Pair, error) {
	var pair Pair
	var err error
	pair.Private, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}
	return &pair, nil
}

type RSATransport struct {
	N string `json:"n"`
	E int    `json:"e"`
	M []byte `json:"m"`
	A string `json:"a"`
}

func (p *Pair) NewRSATransport() *RSATransport {
	public := &p.Private.PublicKey
	var ret RSATransport
	ret.E = public.E
	ret.N = public.N.String()
	return &ret
}

func (rt *RSATransport) Covert() *rsa.PublicKey {
	var ret rsa.PublicKey
	ret.E = rt.E
	ret.N = big.NewInt(0)
	ret.N.SetString(rt.N, 10)
	return &ret
}

func (rt *RSATransport) Encrypt(msg []byte) error {
	var err error
	public := rt.Covert()
	rt.A = "cyt"
	rt.M, err = rsa.EncryptOAEP(md5.New(), rand.Reader, public, msg, nil)
	return err
}

func (rt *RSATransport) Decrypt(privateKey *rsa.PrivateKey) ([]byte, error) {
	msg, err := rsa.DecryptOAEP(md5.New(), rand.Reader, privateKey, rt.M, nil)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
