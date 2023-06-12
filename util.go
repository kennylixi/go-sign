package sign

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	// 保证每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
}

// RandString 生成随机字符串
func RandString(n int) string {
	var count int
	if n%2 == 1 {
		count = (n + 1) / 2
	} else {
		count = n / 2
	}
	result := make([]byte, count)
	rand.Read(result)
	return hex.EncodeToString(result)[:n]
}
