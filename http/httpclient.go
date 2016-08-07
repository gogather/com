package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/gogather/com"
	"github.com/gogather/com/log"
)

// Jar - cookie jar
type Jar struct {
	lk          sync.Mutex
	CookiesData map[string][]*http.Cookie
}

// NewJar - new a Jar
func NewJar() *Jar {
	jar := new(Jar)
	jar.CookiesData = make(map[string][]*http.Cookie)
	return jar
}

func (j *Jar) addCookie(host string, cookie *http.Cookie) {
	hostCookiesData, ok := j.CookiesData[host]
	if ok {
		finded := false
		for i := 0; i < len(hostCookiesData); i++ {
			c := hostCookiesData[i]
			if c.Name == cookie.Name {
				hostCookiesData[i] = cookie
				finded = true
			}
		}

		if !finded {
			hostCookiesData = append(hostCookiesData, cookie)
		}

		j.CookiesData[host] = hostCookiesData
	} else {
		j.CookiesData[host] = append(hostCookiesData, cookie)
	}

}

// SetCookies - set cookies
func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.lk.Lock()
	for i := 0; i < len(cookies); i++ {
		j.addCookie(u.Host, cookies[i])
	}
	j.lk.Unlock()
}

// Cookies - get cookie
func (j *Jar) Cookies(u *url.URL) []*http.Cookie {
	return j.CookiesData[u.Host]
}

// HTTPClient - http client
type HTTPClient struct {
	cookiePath string
	jar        *Jar
	client     *http.Client
}

// NewHTTPClient - new an HTTPClient from cookiePath
func NewHTTPClient(cookiePath string) *HTTPClient {
	hc := &HTTPClient{}
	hc.cookiePath = cookiePath
	jar := NewJar()
	jsonData, err := com.ReadFileString(cookiePath)
	if err == nil {
		err = json.Unmarshal([]byte(jsonData), jar)
		if err != nil {
			log.Warnln("illeage cookies jar file")
		}
	}

	hc.jar = jar

	hc.client = &http.Client{Transport: nil, CheckRedirect: nil, Jar: hc.jar, Timeout: 0}

	return hc
}

func (h *HTTPClient) serialze() {
	jar := h.jar
	jsonData, err := com.JsonEncode(jar)
	if err == nil {
		com.WriteFile(h.cookiePath, string(jsonData))
	}
}

// Post - post method
func (h *HTTPClient) Post(urlstr string, parm url.Values) (string, error) {
	resp, err := h.client.PostForm(urlstr, parm)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	h.serialze()
	return string(b), err
}

// Get - get method
func (h *HTTPClient) Get(urlstr string) (string, error) {
	resp, err := h.client.Get(urlstr)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	h.serialze()
	return string(b), err
}
