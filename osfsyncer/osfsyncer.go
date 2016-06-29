//Package osfsyncer implements the sync.Filesyncer interface based on the os package
package osfsyncer

import (
	"io"
	"os"
	"path/filepath"
)

//Osfsyncer implements the sync.Filesyncer interface
type Osfsyncer struct {
}

func (fs *Osfsyncer) Copy(src, dst string) error {
	err := os.MkdirAll(filepath.Dir(dst), 0777)
	if err != nil {
		return err
	}
	fdst, err := os.Create(dst)
	defer fdst.Close()
	if err != nil {
		return err
	}

	fsrc, err := os.Open(src)
	defer fsrc.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(fdst, fsrc)
	if err != nil {
		return err
	}
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime())
	if err != nil {
		return err
	}
	return nil
}

func (fs *Osfsyncer) GetNewerFile(file1, file2 string) (string, error) {
	f1Info, err := os.Stat(file1)
	if err != nil {
		return "", err
	}
	f2Info, err := os.Stat(file2)
	if err != nil {
		return "", err
	}
	if f2Info.ModTime().After(f1Info.ModTime()) {
		return file2, nil
	}
	return file1, nil
}