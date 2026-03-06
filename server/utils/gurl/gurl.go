package gurl

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"hz-admin-base/global"
	"hz-admin-base/utils/convert"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/proxy"
)

type Gurl struct {
	Url             string
	Data            any               // post数据
	BodyType        BodyType          // post数据类型
	Param           map[string]string // query参数
	Method          string
	Header          map[string]string
	setCookie       map[string]string       // 请求头的cookie
	setCookieByHttp map[string]*http.Cookie // 请求头的cookie

	Response   *http.Response // 响应结果
	option     Option
	debug      *Debug
	client     *http.Client //custom client
	sock5Proxy Sock5Proxy
}

type Option struct {
	Timeout    time.Duration // request timeout
	Proxy      string
	SkipVerify bool // skip ssl verify

	Before func(*Gurl)
	After  func(*Gurl)
}

type Sock5Proxy struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Debug struct {
	Enable    bool
	HideBody  bool
	Skip      int
	Payload   []byte // post数据的字符串形式
	ResBody   []byte // 响应数据体
	Err       error  // 错误信息
	Consuming int64  // 请求耗时(毫秒)
}

func (d *Debug) set(g *Gurl, body []byte, err error) {
	d.ResBody = body
	d.Err = err

	if d.HideBody {
		global.SkipLog(g.debug.Skip + 2).Info(fmt.Sprintf(
			"\ngurl Do data {url: %s %s} %dμs \n{request: %s}",
			g.Method, g.Url, d.Consuming, d.Payload))
	} else {
		global.SkipLog(g.debug.Skip + 2).Info(fmt.Sprintf(
			"\ngurl Do data {url: %s %s} %dμs \n{request: %s}\n{response: %s}",
			g.Method, g.Url, d.Consuming, d.Payload, string(body)))
	}
}

type BodyType int

const (
	TEXT BodyType = iota + 1
	FORM
	JSON
	XML
)

func NewGurl(method, url string) *Gurl {
	return &Gurl{Url: url, Method: method, debug: &Debug{}, client: &http.Client{}}
}

// Set options
func (g *Gurl) Set(option Option) *Gurl {
	g.option = option
	return g
}

func (g *Gurl) Sock5Proxy(s Sock5Proxy) *Gurl {
	g.sock5Proxy = s
	return g
}

func (g *Gurl) Debug(enable ...bool) *Gurl {
	if len(enable) == 0 {
		g.debug.Enable = false
	} else {
		g.debug.Enable = enable[0]
	}
	return g
}

func (g *Gurl) DebugHideBody(enable bool) *Gurl {
	g.debug.HideBody = enable
	return g
}

func (g *Gurl) DebugSkip(skip int) *Gurl {
	g.debug.Skip = skip
	return g
}

func (g *Gurl) GetDebug() *Debug {
	return g.debug
}

// SetData Data Set data and the type, default is JSON
func (g *Gurl) SetData(data any, body ...BodyType) *Gurl {
	g.Data = data
	if len(body) > 0 {
		g.BodyType = body[0]
	} else {
		g.BodyType = JSON
	}
	return g
}

// SetParam Param Set params
func (g *Gurl) SetParam(param map[string]string) *Gurl {
	g.Param = param
	return g
}

// SetHeader Header Set headers
func (g *Gurl) SetHeader(header map[string]any) *Gurl {
	headers := map[string]string{}
	for k, v := range header {
		headers[k] = convert.StringMust(v)
	}
	g.Header = headers
	return g
}

// SetCookie setCookie Set cookies
func (g *Gurl) SetCookie(cookie map[string]string) *Gurl {
	g.setCookie = cookie
	return g
}

func (g *Gurl) SetCookieByHttp(cookie map[string]*http.Cookie) *Gurl {
	g.setCookieByHttp = cookie
	return g
}

func (g *Gurl) GetCookieMap() map[string]string {
	m := make(map[string]string)
	for _, cookie := range g.Response.Cookies() {
		m[cookie.Name] = cookie.Value
	}
	return m
}

func (g *Gurl) GetCookieMapByHttp() map[string]*http.Cookie {
	m := make(map[string]*http.Cookie)
	for _, cookie := range g.Response.Cookies() {
		m[cookie.Name] = cookie
	}
	return m
}

// Client Set custom client
func (g *Gurl) Client(client *http.Client) *Gurl {
	g.client = client
	return g
}

// Combined urls and parameters
func (g *Gurl) urlWithParam() (err error) {
	if g.Param == nil {
		return
	}

	var u *url.URL
	if u, err = url.Parse(g.Url); err != nil {
		return
	}

	q := u.Query()
	for k, v := range g.Param {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	g.Url = u.String()

	return
}

// setupSock5Proxy 设置Sock5代理
func (g *Gurl) setupSock5Proxy(tr *http.Transport) error {
	// 构建代理地址
	proxyAddr := net.JoinHostPort(g.sock5Proxy.Host, strconv.Itoa(g.sock5Proxy.Port))

	// 创建SOCKS5代理拨号器
	var (
		dialer proxy.Dialer
		auth   *proxy.Auth
		err    error
	)
	if g.sock5Proxy.User != "" && g.sock5Proxy.Password != "" {
		// 带认证的SOCKS5代理
		auth = &proxy.Auth{
			User:     g.sock5Proxy.User,
			Password: g.sock5Proxy.Password,
		}
	}

	dialer, err = proxy.SOCKS5("tcp", proxyAddr, auth, proxy.Direct)
	if err != nil {
		return err
	}

	// 设置自定义拨号函数
	tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}

	return nil
}

func (g *Gurl) Request() (err error) {
	if g.Url == "" {
		return errors.New("no url")
	}
	if err = g.urlWithParam(); err != nil {
		return
	}

	if g.Method == "" {
		return errors.New("no method")
	}
	g.Method = strings.ToUpper(g.Method)

	var payload []byte

	if g.Data != nil && g.Method != "GET" {
		switch g.BodyType {
		case TEXT:
			payload = []byte(g.Data.(string))
		case FORM:
			data, ok := g.Data.(map[string]any)
			if !ok {
				return
			}
			formData := url.Values{}
			for k, v := range data {
				formData.Set(k, convert.StringMust(v))
			}
			payload = []byte(strings.ReplaceAll(formData.Encode(), "+", "%20"))
		case JSON:
			if payload, err = global.Json.Marshal(g.Data); err != nil {
				return err
			}
		case XML:
			if payload, err = xml.Marshal(g.Data); err != nil {
				return err
			}
		}
	}

	if g.debug.Enable {
		g.debug.Payload = payload
	}

	request, err := http.NewRequest(g.Method, g.Url, bytes.NewReader(payload))
	if err != nil {
		return
	}

	switch g.BodyType {
	case FORM:
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case JSON:
		request.Header.Set("Content-Type", "application/json")
	}

	if g.Header != nil {
		for k, v := range g.Header {
			request.Header.Set(k, v)
		}
	}

	if g.setCookie != nil {
		for k, v := range g.setCookie {
			request.AddCookie(&http.Cookie{Name: k, Value: v})
		}
	}

	if g.setCookieByHttp != nil {
		for _, v := range g.setCookieByHttp {
			request.AddCookie(v)
		}
	}

	// options
	opt := g.option

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: opt.SkipVerify},
	}

	//使用代理
	if opt.Proxy != "" {
		proxy, err := url.Parse(opt.Proxy)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	// 使用Sock5代理
	if g.sock5Proxy.Host != "" && g.sock5Proxy.Port > 0 {
		if err := g.setupSock5Proxy(tr); err != nil {
			return fmt.Errorf("failed to setup socks5 proxy: %v", err)
		}
	}

	g.client.Transport = tr
	g.client.Timeout = opt.Timeout

	g.Response, err = g.client.Do(request)
	return err
}

func (g *Gurl) Do() ([]byte, error) {
	var (
		s    = time.Now()
		body []byte
	)
	err := g.Request()
	if err != nil {
		return nil, err
	}
	defer g.Response.Body.Close()

	if g.debug.Enable {
		g.debug.Consuming = time.Since(s).Microseconds()
	}

	if g.Response.StatusCode != 200 {
		err = errors.New(g.Response.Status)
	} else {
		body, err = io.ReadAll(g.Response.Body)
	}

	// 未开启调试模式直接返回
	if g.debug.Enable {
		g.debug.set(g, body, err)
	}
	return body, err
}

func (g *Gurl) Decode(data any) error {
	var (
		s    = time.Now()
		body []byte
	)
	err := g.Request()
	if err != nil {
		return err
	}
	defer g.Response.Body.Close()

	if g.debug.Enable {
		g.debug.Consuming = time.Since(s).Milliseconds()
	}

	if g.Response.StatusCode != 200 {
		err = errors.New(g.Response.Status)
	} else {
		if g.debug.Enable {
			// debug时直接读取响应体内容
			if body, err = io.ReadAll(g.Response.Body); err != nil {
				return err
			}
			err = global.Json.Unmarshal(body, data)
		} else {
			// 非调试模式直接流式解码
			err = global.Json.NewDecoder(g.Response.Body).Decode(data)
		}
	}

	// 未开启调试模式直接返回
	if g.debug.Enable {
		g.debug.set(g, body, err)
	}
	return err
}

func (g *Gurl) Send() error {
	var (
		s = time.Now()
	)
	err := g.Request()
	if err != nil {
		return err
	}
	defer g.Response.Body.Close()

	if g.debug.Enable {
		g.debug.Consuming = time.Since(s).Microseconds()
	}

	if g.Response.StatusCode != 200 {
		err = errors.New(g.Response.Status)
	} else {
		if g.debug.Enable {
			var body []byte
			if body, err = io.ReadAll(g.Response.Body); err != nil {
				return err
			}
			_ = body
		}
	}
	// 未开启调试模式直接返回
	if g.debug.Enable {
		g.debug.set(g, []byte{}, err)
	}
	return err
}

func (g *Gurl) Html() (*goquery.Document, error) {
	var (
		s = time.Now()
	)
	err := g.Request()
	if err != nil {
		return nil, err
	}
	defer g.Response.Body.Close()

	if g.debug.Enable {
		g.debug.Consuming = time.Since(s).Microseconds()
	}

	var (
		doc *goquery.Document
	)
	if g.Response.StatusCode != 200 {
		err = errors.New(g.Response.Status)
	} else {
		if g.debug.Enable {
			// debug时直接读取响应体内容 方便断点查看
			var body []byte
			if body, err = io.ReadAll(g.Response.Body); err != nil {
				return nil, err
			}
			doc, err = goquery.NewDocumentFromReader(bytes.NewReader(body))
		} else {
			doc, err = goquery.NewDocumentFromReader(g.Response.Body)
		}
	}
	// 未开启调试模式直接返回
	if g.debug.Enable {
		g.debug.set(g, []byte{}, err)
	}

	return doc, err
}
