package json

import (
	"encoding/json"
	"errors"
)

func Encode(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func MustEncode(v any) []byte {
	b, err := Encode(v)
	if err != nil {
		return []byte("")
	}
	return b
}

func Decode(data []byte, v any) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func MustDecode(data []byte, v any) {
	Decode(data, v)
}

func DecodeStr(str string, v any) error {
	if str == "" {
		return errors.New("empty json string")
	}

	if !json.Valid([]byte(str)) {
		return errors.New("invalid json string")
	}

	return Decode([]byte(str), v)
}

func Encode2Str(v any) (string, error) {
	b, err := Encode(v)
	return string(b), err
}

func MustEncode2Str(v any) string {
	b, _ := Encode(v)
	return string(b)
}

func Valid(s string) bool {
	if s == "" {
		return false
	}
	return json.Valid([]byte(s))
}
