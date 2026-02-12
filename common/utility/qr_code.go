package utility

import (
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
)

func GenerateQRCode(data string) (string, error) {
	qrCode, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		return "", err
	}

	scaledImg := resize.Resize(300, 300, img, resize.Lanczos3)

	var buf bytes.Buffer
	if err := png.Encode(&buf, scaledImg); err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}
