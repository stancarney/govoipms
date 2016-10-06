package v1

import (
	"net/http"
	"log"
	"encoding/json"
	"errors"
	"fmt"
)

type Client struct {
	Url      string
	Username string
	Password string
	Debug    bool
}

type StatusResp interface {
	GetStatus() string
}

type BaseResp struct {
	Status string `json:"status"`
}

func (b *BaseResp) GetStatus() string {
	return b.Status
}

type StringValueDescription struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type NumberValueDescription struct {
	Value       json.Number `json:"value"`
	Description string `json:"description"`
}

func (c *Client) Call(req *http.Request, respStruct interface{}) (*http.Response, error) {

	if c.Debug {
		log.Println("URL:", req.URL)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if c.Debug {
		log.Println("Response:", resp)
	}

	//str, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(str))

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(respStruct); err != nil {
		panic(err)
	}

	return resp, err
}

func (c *Client) Get(url string, entity interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.Call(req, entity)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	s, ok := entity.(StatusResp)
	if ok && s.GetStatus() != "success" {
		return errors.New(s.GetStatus())
	}

	return nil
}

func (c *Client) BaseUrl(apiMethod string) string {
	return fmt.Sprintf("%s?api_username=%s&api_password=%s&method=%s", c.Url, c.Username, c.Password, apiMethod)
}

func (c *Client) NewGeneralAPI() *GeneralAPI {
	return &GeneralAPI{c}
}

func (c *Client) NewAccountAPI() *AccountAPI {
	return &AccountAPI{c}
}

