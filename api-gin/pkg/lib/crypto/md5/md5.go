package md5

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// EncodeAny encodes any type of data into a hex string using MD5 hashing.
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

// Encode encodes a string into a hex string using MD5 hashing.
func Encode(str string) string {
	return EncodeBytes([]byte(str))
}

// EncodeBytes encodes a byte slice into a hex string using MD5 hashing.
func EncodeBytes(d []byte) string {
	sum := md5.Sum(d)
	return hex.EncodeToString(sum[:])
}

// Encode16 encodes a string into a hex string using MD5 hashing and returns the middle 16 characters.
func Encode16(str string) string {
	s := Encode(str)
	return s[8:24]
}

// EncodeBytes16 encodes a byte slice into a hex string using MD5 hashing and returns the middle 16 characters.
func EncodeBytes16(d []byte) string {
	s := EncodeBytes(d)
	return s[8:24]
}

// EncodeAny16 encodes any type of data into a hex string using MD5 hashing and returns the middle 16 characters.
func EncodeAny16(d any) string {
	s := EncodeAny(d)
	if s == "" {
		return ""
	}
	return s[8:24]
}
