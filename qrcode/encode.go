package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

func QrEncode() {
	//qrCode, _ := qr.Encode("https://www.jianshu.com/p/abcdew", qr.M, qr.Auto)
	qrCode, _ := qr.Encode(`{"a": "bca", "arr": [1,2,3], "j": {"data": 4, "code":5 }}`, qr.H, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create("./data/image/test.png")
	defer file.Close()
	png.Encode(file, qrCode)
}
