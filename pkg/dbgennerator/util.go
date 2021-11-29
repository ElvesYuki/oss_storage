package dbgennerator

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"unicode"
)

// Camel2Case 驼峰写法转为下划线写法
func Camel2Case(src string) string {
	buffer := NewBuffer()
	for i, r := range src {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(src string) string {
	src = strings.Replace(src, "_", " ", -1)
	src = strings.Title(src)
	return strings.Replace(src, " ", "", -1)
}

// UcFirst 首字母大写
func UcFirst(src string) string {
	for i, v := range src {
		return string(unicode.ToUpper(v)) + src[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(src string) string {
	for i, v := range src {
		return string(unicode.ToLower(v)) + src[i+1:]
	}
	return ""
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
