package goutils

import "encoding/base64"

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(src string) string {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(data)
}
