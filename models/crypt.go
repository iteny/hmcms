package models

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func Sha1(buf []byte) string {
	hash := sha1.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
