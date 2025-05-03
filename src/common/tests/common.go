package tests

import (
	"bytes"
	"log"

	"github.com/AmrSaber/random/src/common"
)

func SetupTest() (*bytes.Buffer, func()) {
	var buf bytes.Buffer
	oldStd := common.Std
	common.Std = log.New(&buf, "", 0)
	return &buf, func() { common.Std = oldStd }
}
