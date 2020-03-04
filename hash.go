package e2pay

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// HastToMD5 is used to encrypt to md5
func HastToMD5(str string) string {
	data := []byte(str)
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum(data)))
}
