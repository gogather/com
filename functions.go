package com

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

const (
	_VERSION = "0.1.0307"
)

func Version() string {
	return _VERSION
}

// 获取用户头像
func GetGravatar(email string) string {
	return "http://www.gravatar.com/avatar/" + Md5(strings.ToUpper(email))
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

// 四舍五入
func Round(inputValue float64, places int) float64 {
	var result float64
	scaleFactor := math.Pow10(places)
	scaledValue := inputValue * scaleFactor
	if math.IsInf(scaledValue, 0) || math.IsNaN(scaledValue) {
		return inputValue
	}
	if scaledValue >= 0.0 {
		result = math.Ceil(scaledValue)
		if (result - scaledValue) > 0.50000000001 {
			result -= 1.0
		}
	} else {
		result = math.Ceil(-scaledValue)
		if (result + scaledValue) > 0.50000000001 {
			result -= 1.0
		}
		result = -result
	}
	scaledValue = result / scaleFactor

	if !math.IsInf(scaledValue, 0) {
		return scaledValue
	}

	return result
}

// http GET
func Get(requestURL string) (string, error) {
	response, err := http.Get(requestURL)
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
