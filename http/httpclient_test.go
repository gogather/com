package http

import (
	"testing"

	"github.com/gogather/com/log"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_HTTPClient(t *testing.T) {
	Convey("Http client test", t, func() {
		client := NewProxyHTTPClient("socks5://127.0.0.1:1080")
		content, _ := client.Get("https://www.google.com.hk")
		log.Println(content)
	})

}
