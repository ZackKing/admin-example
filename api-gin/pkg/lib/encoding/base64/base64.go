package base64

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

func Encode(str string) string {
	return EncodeBytes([]byte(str))
}

func EncodeBytes(b []byte) string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	return string(dst)
}

func EncodeAny(d any) string {
	switch d := d.(type) {
	case string:
		return Encode(d)
	case []byte:
		return EncodeBytes(d)
	default:
		ds, err := json.Marshal(d)
		if err != nil {
			return ""
		}
		return EncodeBytes(ds)
	}
}

func EncodeFile(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return EncodeBytes(b), nil
}
