package e2pay

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Client is ..
type Client struct {
	Host      string
	ClientID  string
	SecretKey string
	SourceID  string
	PartnerID string
	LogLevel  int
	Logger    *log.Logger
}

var (
	defHTTPTimeout = 15 * time.Second
	httpClient     = &http.Client{Timeout: defHTTPTimeout}
)

// NewClient is ...
func NewClient() Client {
	return Client{
		LogLevel: 2,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

// NewRequest is ...
func (c *Client) NewRequest(method, fullPath string, headers map[string]string, body io.Reader) (*http.Request, error) {
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Error create new request: ", err)
		}
		return nil, err
	}
	if headers != nil {
		for k, vv := range headers {
			req.Header.Set(k, vv)
		}
	}
	return req, nil
}

// ExecuteRequest is ...
func (c *Client) ExecuteRequest(req *http.Request, v interface{}) error {
	logLevel := c.LogLevel
	logger := c.Logger
	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}
	start := time.Now()
	res, err := httpClient.Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("e2pay response: ", string(resBody))
	}

	if v != nil && res.StatusCode == 200 {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}
	}
	return nil
}

// Call is ...
func (c *Client) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, header, body)
	if err != nil {
		return err
	}
	return c.ExecuteRequest(req, v)
}
