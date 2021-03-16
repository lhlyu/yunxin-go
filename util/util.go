package util

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

// SHA1toStr 使用 SHA1 进行编码
func SHA1toStr(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func ToJsonStr(v interface{}) string {
	if v == nil {
		return ""
	}
	b, _ := json.Marshal(v)
	return string(b)
}
