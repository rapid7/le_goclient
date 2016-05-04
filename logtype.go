package logentries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type LogType struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Key         string `json:"key"`
	Shortcut    string `json:"shortcut"`
	ApiObject
}

type LogTypeClient struct {
	client
}

type LogTypeListRequest struct {
}

type LogTypeListResponse struct {
	List []LogType
	ApiResponse
}

func (u *LogTypeClient) ReadDefault(defaultLogTypeListRequest LogTypeListRequest) ([]LogType, error) {
	return u.read("list_logtypes_default", defaultLogTypeListRequest)
}

func (u *LogTypeClient) Read(defaultLogTypeListRequest LogTypeListRequest) ([]LogType, error) {
	return u.read("list_logtypes", defaultLogTypeListRequest)
}

func (u *LogTypeClient) read(requestType string, logTypeListRequest LogTypeListRequest) ([]LogType, error) {
	form := url.Values{}
	form.Add("request", requestType)
	form.Add("id", "terraform")
	resp, err := u.PostForm(form)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var response LogTypeListResponse
		json.NewDecoder(resp.Body).Decode(&response)
		if response.Response == "ok" {
			return response.List, nil
		} else {
			return nil, fmt.Errorf("failed to retrieve default log type list: %s", response.ResponseReason)
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return nil, fmt.Errorf("Could not retrieve default log set info: %s", string(body))
}

func NewLogTypeClient(account_key string) *LogTypeClient {
	return &LogTypeClient{defaultClient(account_key)}
}
