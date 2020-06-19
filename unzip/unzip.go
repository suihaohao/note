package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println("read zip file err:", err)
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				fmt.Println("open file err:", err)
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				fmt.Println("copy  file err:", err)
				return err
			}
		}
	}
	return nil
}
