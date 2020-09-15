package awair

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	baseAwairUrl = "https://developer-apis.awair.is/"
	userAgent    = "go-awair"
)

type Client struct {
	client *http.Client

	baseUrl *url.URL

	common    service
	userAgent string

	User *UserService
}

type service struct {
	client *Client
}

func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	fullUrl, err := c.baseUrl.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	req, err := http.NewRequest(method, fullUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		client: httpClient,
	}

	c.baseUrl, _ = url.Parse(baseAwairUrl)
	c.common.client = c
	c.userAgent = userAgent

	c.User = (*UserService)(&c.common)

	return c
}
