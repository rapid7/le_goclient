package logentries

import (
    "net/http"
    "net/url"
    "strings"
    "encoding/json"
    "io/ioutil"
    "fmt"
)

type AccountClient struct {
    AccountKey string
}

type AccountReadRequest struct {
}

type AccountReadResponse struct {
    User
    ApiResponse
}

func (u *AccountClient) Read(readRequest AccountReadRequest) (*AccountReadResponse, error) {
    form := url.Values{}
    form.Add("request", "get_user")
    form.Add("load_hosts", "1")
    form.Add("load_logs", "1")
    form.Add("load_alerts", "0")
    form.Add("user_key", u.AccountKey)
    form.Add("id", "terraform")
    resp, err := http.Post(
        "https://api.logentries.com/", 
        "application/x-www-form-urlencoded", 
        strings.NewReader(form.Encode()))

    if err != nil {
        return nil, err
    }

    if resp.StatusCode == 200 {
        var response AccountReadResponse
        json.NewDecoder(resp.Body).Decode(&response)
        return &response, nil
    }

    body, _ := ioutil.ReadAll(resp.Body)
    return nil, fmt.Errorf("Could not retrieve account info: %s", string(body))
}

func NewAccountClient(account_key string) *AccountClient {
    account := AccountClient{AccountKey: account_key}
    return &account
}