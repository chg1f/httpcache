package httpcache

import (
	"io"
	netHttp "net/http"
	netUrl "net/url"
	"strings"
)

type Client struct {
	netHttp.Client
	Cache
}

func (c *Client) Do(req *netHttp.Request, opts ...func(*CacheOption)) (resp *netHttp.Response, err error) {
	if len(opts) > 0 {
		opt := new(CacheOption)
		for ix := range opts {
			opts[ix](opt)
		}
		if opt.Key != "" {
			if resp, ok := c.Cache.Get(opt); ok {
				return resp, nil
			}
			defer func() {
				if err != nil {
					_ = c.Cache.Set(opt, resp)
				}
			}()
		}
	}
	resp, err = c.Client.Do(req)
	return
}
func (c *Client) Get(url string, opts ...func(*CacheOption)) (resp *netHttp.Response, err error) {
	req, err := netHttp.NewRequest(netHttp.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
func (c *Client) Head(url string, opts ...func(*CacheOption)) (resp *netHttp.Response, err error) {
	req, err := netHttp.NewRequest(netHttp.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, opts...)
}
func (c *Client) Post(url, contentType string, body io.Reader, opts ...func(*CacheOption)) (resp *netHttp.Response, err error) {
	req, err := netHttp.NewRequest(netHttp.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req, opts...)
}
func (c *Client) PostForm(url string, data netUrl.Values, opts ...func(*CacheOption)) (resp *netHttp.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), opts...)
}
