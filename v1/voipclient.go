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
	"bytes"
)

type VOIPClient struct {
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

func NewVOIPClient(url, username, password string, debug bool) *VOIPClient {
	return &VOIPClient{url, username, password, debug}
}

func (c *VOIPClient) Call(req *http.Request, respStruct interface{}) (*http.Response, error) {

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

func (c *VOIPClient) Get(url string, entity interface{}) error {
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

func (c *VOIPClient) Post(method string, entity interface{}, respStruct interface{}) error {
	
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("api_username", c.Username)
	bodyWriter.WriteField("api_password", c.Password)
	bodyWriter.WriteField("method", method)

	if err := WriteStruct(bodyWriter, entity); err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("POST", c.URL, bodyBuf)
	req.Header.Set("Content-Type", contentType)
	resp, err := c.Call(req, respStruct)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	s, ok := respStruct.(StatusResp)
	if ok && s.GetStatus() != "success" {
		return errors.New(s.GetStatus())
	}
	
	return nil
}

func (c *VOIPClient) BaseUrl(apiMethod string) string {
	return fmt.Sprintf("%s?api_username=%s&api_password=%s&method=%s", c.URL, c.Username, c.Password, apiMethod)
}

func (c *VOIPClient) NewGeneralAPI() *GeneralAPI {
	return &GeneralAPI{c}
}

func (c *VOIPClient) NewAccountsAPI() *AccountsAPI {
	return &AccountsAPI{c}
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

