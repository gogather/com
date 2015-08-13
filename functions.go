package com

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 创建GUID
func CreateGUID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

// 截取子串
func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	return string(rs[begin:end])
}

// 创建随机串
func RandString(n int) string {
	guid := CreateGUID()
	return SubString(guid, 0, n)
}

// 检查用户名
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

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// 获取用户头像
func GetGravatar(email string) string {
	return "http://www.gravatar.com/avatar/" + Md5(strings.ToUpper(email))
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func PathExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 读取文件
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// 读取文本文件
func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return string(fd)
}

// 切割关键词为html片段
func TagSplit(keywords string) string {
	if "" == keywords {
		return ""
	}

	content := ""
	tags := strings.Split(keywords, ",")
	for _, value := range tags {
		content = content + fmt.Sprintf(`<a class="tags" href="/tag/%s/1">%s</a>,`, value, value)
	}
	return content
}

// json编码
func JsonEncode(data interface{}) (string, error) {
	a, err := json.Marshal(data)
	return string(a), err
}

// json解码
func JsonDecode(data string) (interface{}, error) {
	dataByte := []byte(data)
	var dat interface{}

	err := json.Unmarshal(dataByte, &dat)
	return dat, err
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 字符串写入文件
func WriteFile(fullpath string, str string) error {
	data := []byte(str)
	return ioutil.WriteFile(fullpath, data, 0644)
}

// 创建文件夹
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// 四舍五入
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

// http GET
func Get(reqUrl string) (string, error) {
	response, err := http.Get(reqUrl)
	if nil != err {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		response.Body.Close()
		return "", err
	}
	return string(body), nil
}

func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

// Copyfile
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func Unicode(rs string) string {
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

func HTMLEncode(rs string) string {
	html := ""
	for _, r := range rs {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}
