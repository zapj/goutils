package fileutils

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/zapj/goutils"
)

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func CopyFileExistsSkip(src, dst string) error {
	if IsFile(dst) {
		return os.ErrExist
	}
	return CopyFile(src, dst)
}

func CopyDir(src string, dst string, args ...bool) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo
	skipExistsFile := goutils.LeftOrRight[bool](len(args) == 1, args[0], false)
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				return err
			}
		} else {
			if skipExistsFile {
				if err = CopyFileExistsSkip(srcfp, dstfp); err != nil {
					return err
				}

			} else if err = CopyFile(srcfp, dstfp); err != nil {
				return err
			}
		}
	}
	return nil
}
