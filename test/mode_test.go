package test

import (
	"common"
	"testing"
)

func TestMode(t *testing.T) {
	mode := common.Mode(1)
	println(mode.String())
	println(common.Modes[mode.String()])

	mode = common.Mode(10086)
	println(mode.String())
}
