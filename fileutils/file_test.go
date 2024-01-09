package fileutils_test

import (
	"testing"

	"github.com/zapj/goutils/fileutils"
)

func Test_FilePwd(t *testing.T) {
	t.Log(fileutils.Pwd())
}

func Test_TempFileName(t *testing.T) {
	t.Log(fileutils.TempFileName("zap_"))
}
