package v1

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"errors"
	"mime/multipart"
	"reflect"
	"strings"
	"net/http/httputil"
)

type Client struct {
	URL      string
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

func NewClient(url, username, password string, debug bool) *Client {
	return &Client{url, username, password, debug}
}

func (c *Client) Call(req *http.Request, respStruct interface{}) (*http.Response, error) {

	//req.Header.Set("Content-Type", "multipart/form-data")
	if c.Debug {
		out, _ := httputil.DumpRequest(req, true)
		log.Println(string(out))
	}

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
		return nil, err
	}

	return resp, err
}

func (c *Client) Get(url string, entity interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
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
	return fmt.Sprintf("%s?api_username=%s&api_password=%s&method=%s", c.URL, c.Username, c.Password, apiMethod)
}

func (c *Client) NewGeneralAPI() *GeneralAPI {
	return &GeneralAPI{c}
}

func (c *Client) NewAccountAPI() *AccountAPI {
	return &AccountAPI{c}
}

func WriteStruct(writer *multipart.Writer, i interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(i))

	for i := 0; i < val.NumField(); i++ {

		structField := val.Type().Field(i)

		name := strings.TrimSuffix(structField.Tag.Get("json"), ",omitempty") //TODO:Stan the omitempty is rather fragile.
		if name == "" {
			name = strings.ToLower(structField.Name)
		}

		value := val.Field(i).Interface().(string)

		if err := writer.WriteField(name, value); err != nil {
			return err
		}
	}

	return nil
}
