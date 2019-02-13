package topdeliveryRestClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-topdelivery-rest-client/utils"
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

func NewClient(baseUrl string, login string, password string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		config: config{
			BaseUrl:  baseUrl,
			Login:    login,
			Password: password,
		},
	}
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

	defer resp.Body.Close();

	if resp.StatusCode != http.StatusOK {
		var e topdeliveryErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&e);
		return errors.New("bad response " + err.Error() + e.Error + e.Message)

	}

	var f topdeliverySuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&f);
	if err != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New("error unmarshaling " + err.Error() + string(bodyBytes))
	}

	err = json.Unmarshal(f.Data.Rows, &bar);
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
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(c.config.BaseUrl+"/Auth", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		return errors.New("error auth request " + err.Error())
	}

	defer resp.Body.Close();

	if resp.StatusCode != http.StatusOK {
		var e topdeliveryErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&e);
		return errors.New("bad auth response " + err.Error() + e.Error + e.Message)
	}

	var f topdeliverySuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&f);
	if err != nil {
		return err
	}

	var rows authData
	err = json.Unmarshal(f.Data.Rows, &rows);
	if err != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New("in auth response topdelivery mistake token. Response: " + string(bodyBytes))
	}

	token := Token{}
	if err := utils.Decode(rows.Token, &token); err != nil {
		return errors.New("error token decoding: " + err.Error())
	}

	c.expiredAt = token.Exp
	c.token = rows.Token

	return nil
}
