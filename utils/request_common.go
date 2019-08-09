package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTPArgs http request args
type HTTPArgs struct {
	Method  string
	Values  map[string]string
	URL     string
	Body    io.Reader
	Headers map[string]string
}

// SetRequestURL set real request url
func (args *HTTPArgs) SetRequestURL() (err error) {
	if args.URL == "" {
		err = fmt.Errorf("url not found: %v", args.URL)
		return
	}
	urlValues := url.Values{}
	for key, value := range args.Values {
		urlValues.Add(key, value)
	}

	urlArgs := urlValues.Encode()
	args.URL = fmt.Sprintf("%s?%s", args.URL, urlArgs)
	return
}

// GetHTTPResponseBytes get url response body bytes
func (args *HTTPArgs) GetHTTPResponseBytes() (content []byte, err error) {
	err = args.SetRequestURL()
	if err != nil {
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest(args.Method, args.URL, args.Body)
	if err != nil {
		return
	}

	for key, value := range args.Headers {
		request.Header.Add(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}

	body := response.Body
	// defer response.Body.Close()
	defer func() {
		if fErr := response.Body.Close(); fErr != nil {
			err = fErr
		}
	}()

	if err != nil {
		return
	}

	content, err = ioutil.ReadAll(body)
	return
}

// APIResponseMeta api response meta
func APIResponseMeta(c *gin.Context, code int, data interface{}) (int, map[string]interface{}) {
	requestID := c.MustGet("RequestID")
	if requestID == "" {
		requestID = "InValidID"
	}

	codeMsg := http.StatusText(code)
	if codeMsg == "" {
		codeMsg = GetStatusText(code)
	}
	return code, gin.H{
		"meta": map[string]interface{}{
			"requestID":   requestID,
			"requestTime": time.Now().Format("2006/01/02 15:04:05"),
		},
		"data": data,
		"status": map[string]interface{}{
			"code": code,
			"msg":  codeMsg,
		},
	}
}
