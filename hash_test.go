package goutils_test

import (
	"testing"

	"github.com/zapj/goutils"
)

func TestMD5(t *testing.T) {
	t.Error(goutils.MD5("test"))
}
