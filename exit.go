package common

import (
	"os"
	"time"
)

func Exit(code Code, fs ...func()) {
	for _, f := range fs {
		f()
	}
	time.Sleep(time.Second * 5)
	os.Exit(int(code))
}
