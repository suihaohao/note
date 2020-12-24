package qrcode

import (
	"fmt"
	"github.com/liyue201/goqr"
	"github.com/tuotoo/qrcode"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func qrd2(imgpath string) (string, error) {
	//use github.com/tuotoo/qrcode
	fi, err := os.Open(imgpath)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	t, err := qrcode.Decode(fi)
	if err != nil {
		return "", err
	}
	return t.Content, nil
}

func qrd1(imagepath string) (string, error) {
	//use github.com/liyue201/goqr
	file, err := os.OpenFile(imagepath, 0, os.ModePerm)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return "", err
	}
	text := ""
	for _, qrCode := range qrCodes {
		text = fmt.Sprintf("%s%s", text, qrCode.Payload)
	}
	return text, nil
}

func QrDecode(imagepath string) (string, error) {
	text, err := qrd1(imagepath)
	if err == nil {
		return text, err
	}
	return qrd2(imagepath)
}
