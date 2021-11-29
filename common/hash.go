package common

import (
	"crypto/md5"
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
