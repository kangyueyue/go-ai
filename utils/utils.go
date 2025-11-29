package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// GetRandomNumbers 生成随机name
func GetRandomNumbers(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < num; i++ {
		// 0-9 随机数
		digit := r.Intn(10)
		code += strconv.Itoa(digit)
	}
	return code
}

// MD5 加密
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
