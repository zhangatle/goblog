package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"
	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/")
	assert.NoError(t, err, "有错误发生， err不为空")
	assert.Equal(t, 200, resp.StatusCode, "应该返回状态码200")
}

func TestAboutPage(t *testing.T) {
	baseURL := "http://localhost:3000"

	// 1. 请求 —— 模拟用户访问浏览器
	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/about")

	// 2. 检测 —— 是否无错误且 200
	assert.NoError(t, err, "有错误发生，err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}

func TestAllPages(t *testing.T) {

	baseURL := "http://localhost:3000"

	// 1. 声明加初始化测试数据
	var tests = []struct {
		method   string
		url      string
		expected int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 404},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/1", 200},
		{"GET", "/articles/1/edit", 200},
		{"POST", "/articles/1", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/333/delete", 404},
	}

	// 2. 遍历所有测试
	for _, test := range tests {
		t.Logf("当前请求 URL: %v \n", test.url)
		var (
			resp *http.Response
			err  error
		)
		// 2.1 请求以获取响应
		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL + test.url)
		}
		// 2.2 断言
		assert.NoError(t, err, "请求 "+test.url+" 时报错")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+" 应返回状态码 "+strconv.Itoa(test.expected))
	}
}
