package http

import (
	"testing"

	"io/ioutil"
	"net/http"

	"github.com/gogather/com/log"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_HTTPClient(t *testing.T) {
	client := NewProxyHTTPClient("socks5://127.0.0.1:1080")
	Convey("Http client test", t, func() {
		content, _ := client.Get("https://www.google.com.hk")
		log.Println(content)
	})

	Convey("Http download test", t, func() {
		client.Download("https://www.google.com.hk", "index.html")
	})

	Convey("Http do", t, func() {
		r, _ := http.NewRequest("GET", "https://www.google.com.hk", nil)
		r.Header.Add("Content-Type", "charset=UTF-8")
		r.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
		resp, _ := client.Do(r)
		content, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(content))
	})

}
