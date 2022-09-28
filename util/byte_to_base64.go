package util

import "encoding/base64"

func ToBase64(b []byte) string {
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b)
}
