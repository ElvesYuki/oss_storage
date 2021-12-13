package common

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func Md5Util(src io.Reader) (md5Value string, err error) {
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}
	md5Byte := md5.Sum(data)
	return hex.EncodeToString(md5Byte[:]), nil
}

func Md5StringUtil(src string) (md5Value string, err error) {
	md5Byte := md5.Sum([]byte(src))
	return hex.EncodeToString(md5Byte[:]), nil
}

func SHA256StringUtil(src string) (md5Value string, err error) {
	sha256Byte := sha256.Sum256([]byte(src))
	return hex.EncodeToString(sha256Byte[:]), nil
}
