package ggutils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipDir(dir, zipFile string, f func(name string)) error {
	zfile, err := os.Create(zipFile)
	if err != nil {
		return err
	}

	zwriter := zip.NewWriter(zfile)
	defer zwriter.Close()
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			path = strings.Replace(path, "\\", "/", -1)
			var fDest io.Writer
			fDest, err = zwriter.Create(path[len(dir)+1:])
			if err != nil {
				return err
			}

			if f != nil {
				f(path)
			}

			var fSrc *os.File
			fSrc, err = os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()

			var n int64
			n, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
			if n < 0 {

			}
		}
		return nil
	})

	return err
}

func UnzipDir(zipFile, dir string) error {
	r, er := zip.OpenReader(zipFile)
	if er != nil {
		return er
	}
	defer r.Close()

	for _, f := range r.File {
		er = func() error {
			path := dir + string(filepath.Separator) + f.Name
			er = os.MkdirAll(filepath.Dir(path), 0755)
			if er != nil {
				return er
			}

			var fDest *os.File
			fDest, er = os.Create(path)
			if er != nil {
				return er
			}
			defer fDest.Close()

			var fSrc io.ReadCloser
			fSrc, er = f.Open()
			if er != nil {
				return er
			}
			defer fSrc.Close()

			_, er = io.Copy(fDest, fSrc)
			if er != nil {
				return er
			}
			return nil
		}()

		if er != nil {

		}
	}

	return nil
}

func PathExists(path string) (bool, error) {
	_, er := os.Stat(path)
	if er == nil {
		return true, nil
	}

	if os.IsNotExist(er) {
		return false, nil
	}

	return false, er
}
