package util_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/tarokent10/go-sample/imageutil/util"
)

func TestPng2Jpeg(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.png")
	dest := filepath.Join(wd, "testwork", "nanu.png")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	defer func(dest string) {
		if err := delete(dest); err != nil {
			println(err.Error())
		}
	}(dest)
	if err := util.Png2Jpeg(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
}

func copy(srcpath, destpath string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}

func delete(t string) error {
	return os.Remove(t)
}
