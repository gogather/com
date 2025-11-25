package com

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unsafe"
)

// CreateGUID 创建GUID
func CreateGUID() string {
	randomBytes := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(randomBytes)
}

// SubString 截取子串
func SubString(str string, begin, length int) (substr string) {
	runes := []rune(str)
	runeLength := len(runes)

	if begin < 0 {
		begin = 0
	}
	if begin >= runeLength {
		begin = runeLength
	}
	end := begin + length
	if end > runeLength {
		end = runeLength
	}

	return string(runes[begin:end])
}

// RandString 创建随机串
func RandString(n int) string {
	guid := CreateGUID()
	return SubString(guid, 0, n)
}

// CheckUsername 检查用户名
func CheckUsername(username string) bool {
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}

	for i := 0; i < len(username); i++ {
		if !(username[i] == '_' ||
			(username[i] >= '0' && username[i] <= '9') ||
			(username[i] >= 'a' && username[i] <= 'z') ||
			(username[i] >= 'A' && username[i] <= 'Z')) {
			return false
		}
	}

	return true
}

// Md5 MD5加密
func Md5(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(hash.Sum(nil)))
}

// Sha1 对字符串进行sha1 计算
func Sha1(data string) string {
	sha1Hash := sha1.New()
	io.WriteString(sha1Hash, data)
	return fmt.Sprintf("%x", sha1Hash.Sum(nil))
}

// Strim 去除空白符
func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

// Unicode 转 unicode 编码
func Unicode(inputString string) string {
	result := ""
	for _, r := range inputString {
		runeInt := int(r)
		if runeInt < 128 {
			result += string(r)
		} else {
			result += "\\u" + strconv.FormatInt(int64(runeInt), 16)
		}
	}
	return result
}

// HTMLEncode 对字符串进行HTML编码
func HTMLEncode(inputString string) string {
	html := ""
	for _, r := range inputString {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}

// StringToBytes converts string to byte slice without a memory allocation.
// For more details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// BytesToString converts byte slice to string without a memory allocation.
// For more details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func UniqueStrings(input []string) []string {
	uniqueMap := make(map[string]struct{})
	var result []string

	for _, str := range input {
		if _, exists := uniqueMap[str]; !exists {
			uniqueMap[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func ContainsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, str string) []string {
	var result []string
	for _, item := range slice {
		if item != str {
			result = append(result, item)
		}
	}
	return result
}
