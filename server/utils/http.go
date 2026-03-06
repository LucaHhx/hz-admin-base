package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/chromedp/cdproto/network"
)

// Request 模拟带 headers 和 cookies 的 GET 请求
func Request(request *network.Request, cookies []*network.Cookie) ([]byte, error) {
	fmt.Println("[Request]", request.Method, request.URL)

	// 创建 HTTP 请求
	req, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("构建请求失败: %w", err)
	}

	// 设置请求头（仅 string 类型）
	for key, value := range request.Headers {
		if strVal, ok := value.(string); ok {
			req.Header.Set(key, strVal)
		}
	}

	// 构造 Cookie 字符串
	var cookieList []string
	for _, c := range cookies {
		cookieList = append(cookieList, fmt.Sprintf("%s=%s", c.Name, c.Value))
	}
	if len(cookieList) > 0 {
		req.Header.Set("Cookie", strings.Join(cookieList, "; "))
	}

	// 设置默认 UA（如果原请求没有）
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	}

	// 创建客户端并发起请求（默认跟随重定向）
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("[Response] Status Code:", resp.StatusCode)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	return body, nil
}
