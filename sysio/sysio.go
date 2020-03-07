package sysio

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gitamenet/v2ray-core/common/buf"
	"github.com/gitamenet/v2ray-core/common/platform"
)

type FileReaderFunc func(path string) (io.ReadCloser, error)

var NewFileReader FileReaderFunc = func(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func ReadFile(path string) ([]byte, error) {
	reader, err := NewFileReader(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return buf.ReadAllToBytes(reader)
}

func ReadAsset(file string) ([]byte, error) {
	return ReadFile(platform.GetAssetLocation(file))
}

func CopyFile(dst string, src string) error {
	bytes, err := ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, bytes, 0644)
}
