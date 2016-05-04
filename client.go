package logentries

import (
	"net/http"
	"net/url"
)

type ApiResponse struct {
	Response       string `json:"response"`
	ResponseReason string `json:"reason"`
	Worker         string `json:"worker"`
	Id             string `json:"id"`
}

type ApiObject struct {
	Object string `json:"object"`
}

type Client struct {
	Log     *LogClient
	LogSet  *LogSetClient
	User    *UserClient
	LogType *LogTypeClient
}

type client interface {
	PostForm(url.Values) (*http.Response, error)
	Get(string) (*http.Response, error)
}

type realClient struct {
	AccountKey string
	Endpoint   string
}

func (c *realClient) PostForm(form url.Values) (*http.Response, error) {
	form.Add("user_key", c.AccountKey)
	return http.PostForm(c.Endpoint, form)
}

func (c *realClient) Get(path string) (*http.Response, error) {
	return http.Get(c.Endpoint + c.AccountKey + path)
}

func newFullClient(c client) *Client {
	return &Client{
		Log:     &LogClient{c},
		LogSet:  &LogSetClient{c},
		User:    &UserClient{c},
		LogType: &LogTypeClient{c},
	}
}
func defaultClient(account_key string) client {
	return &realClient{
		AccountKey: account_key,
		Endpoint:   "https://api.logentries.com/",
	}
}

func NewClient(account_key string) *Client {
	return newFullClient(defaultClient(account_key))
}
