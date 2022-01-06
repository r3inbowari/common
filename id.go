package common

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func CreateUUID() string {
	return uuid.New().String()
}

func GetMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
