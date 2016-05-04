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

type LogType struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Key         string `json:"key"`
	Shortcut    string `json:"shortcut"`
	ApiObject
}

type Log struct {
	Name      string `json:"name"`
	Created   int64  `json:"created"`
	Key       string `json:"key"`
	Token     string `json:"token"`
	Follow    string `json:"follow"`
	Retention int64  `json:"retention"`
	Source    string `json:"type"`
	Type      string `json:"logtype"`
	Filename  string `json:"filename"`
	ApiObject
}

type LogSet struct {
	Distver  string `json:"distver"`
	C        int64  `json:"c"`
	Name     string `json:"name"`
	Distname string `json:"distname"`
	Location string `json:"hostname"`
	Key      string `json:"key"`
	Logs     []Log
	ApiObject
}

type User struct {
	UserKey string        `json:"user_key"`
	LogSets []LogSet      `json:"hosts"`
	Apps    []interface{} `json:"apps"`
	Logs    []interface{} `json:"logs"`
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

type Client struct {
	Log     *LogClient
	LogSet  *LogSetClient
	User    *UserClient
	LogType *LogTypeClient
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
