package topdeliveryRestClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/computerslong/go-topdelivery-rest-client/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const tokenName = "istopdeliveryru-token"

type (
	topdeliverySuccessResponse struct {
		Status    int
		Error     string
		Timestamp int
		Message   string
		Data      *responseData
	}

	topdeliveryErrorResponse struct {
		Status    int
		Error     string
		Timestamp int
		Message   string
	}

	responseData struct {
		Page      int
		Timestamp int
		Total     int
		Rows      json.RawMessage
	}

	config struct {
		BaseUrl  string
		Login    string
		Password string
	}

	Client struct {
		httpClient *http.Client
		config     config
		token      string
		expiredAt  int32
	}

	Token struct {
		Iis string
		Aud string
		Iat int
		Exp int32
	}

	TopdeliveryUser struct {
		Id   int
		Type string
	}
)

func New(baseUrl string, login string, password string) (*Client, error) {
	return &Client{
		httpClient: http.DefaultClient,
		config: config{
			BaseUrl:  baseUrl,
			Login:    login,
			Password: password,
		},
	}, nil
}

func (c *Client) callApi(req *http.Request, bar interface{}) error {

	if c.expiredAt == 0 || int32(time.Now().Unix()) >= c.expiredAt {
		if err := c.updateToken(); err != nil {
			return errors.New("error updating token " + err.Error())
		}
	}

	req.Header.Set(tokenName, c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = c.handleResponse(resp, bar); err != nil {
		return errors.New("error handling auth response " + err.Error())
	}

	return nil
}

func (c *Client) handleResponse(response *http.Response, bar interface{}) error {
	if response.StatusCode != http.StatusOK {
		var e topdeliveryErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			return errors.New("bad response " + err.Error() + e.Error + e.Message)
		}
		return errors.New("bad response " + e.Error + e.Message)
	}

	var f topdeliverySuccessResponse
	err := json.NewDecoder(response.Body).Decode(&f)
	if err != nil {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return errors.New("error unmarshaling " + err.Error() + string(bodyBytes))
	}

	err = json.Unmarshal(f.Data.Rows, &bar)
	if err != nil {
		return errors.New("error unmarshaling " + err.Error() + string(f.Data.Rows))
	}

	return nil
}

func (c *Client) updateToken() error {

	type authData struct {
		MemberType string
		MemberId   int
		UserId     int
		Title      string
		Token      string `json:"istopdeliveryru-token"`
	}

	values := map[string]string{"login": c.config.Login, "password": c.config.Password}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.config.BaseUrl+"/Auth", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return ErrBadRequest{err, "auth request"}
	}

	defer resp.Body.Close()

	bar := &authData{}
	if err := c.handleResponse(resp, bar); err != nil {
		return ErrBadResponse{err, "handling auth response"}
	}

	var token Token
	if err := utils.Decode(bar.Token, &token); err != nil {
		return ErrBadResponse{err, "token decoding"}
	}

	c.expiredAt = token.Exp
	c.token = bar.Token

	return nil
}
