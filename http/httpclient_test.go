package http

import (
	"testing"

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
		client.Download("https://storage.googleapis.com/golang/go1.8beta1.windows-amd64.msi", "go1.8beta1.windows-amd64.msi")
	})

}
