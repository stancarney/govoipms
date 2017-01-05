package v1

import (
	"net/http"
	"log"
	"encoding/json"
	"errors"
	"mime/multipart"
	"reflect"
	"strings"
	"net/http/httputil"
	"bytes"
	"net/url"
	"io/ioutil"
	"fmt"
	"encoding"
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

	if c.Debug {
		out, _ := httputil.DumpRequest(req, true)
		log.Println(string(out))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if c.Debug {
		log.Println("Response: ", resp)

		b, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}

		log.Println(string(b))

		body = ioutil.NopCloser(bytes.NewReader(b))
	}

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(respStruct); err != nil {
		return nil, err
	}

	return resp, err
}

func (c *VOIPClient) Get(method string, values url.Values, entity interface{}) error {

	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	values.Add("api_username", c.Username)
	values.Add("api_password", c.Password)
	values.Add("method", method)

	u.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	resp, err := c.Call(req, entity)
	if err != nil {
		return err
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

	if err := c.WriteStruct(bodyWriter, entity); err != nil {
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

// Function to simplify calls that only take a single string argument (i.e. an ID) and only return an error on failure, i.e. status != "success"
func (c *VOIPClient) simpleCall(method, argName, argValue string) error {
	values := url.Values{}
	values.Add(argName, argValue)

	rs := &BaseResp{}
	return c.Get(method, values, rs)
}

func (c *VOIPClient) NewGeneralAPI() *GeneralAPI {
	return &GeneralAPI{c}
}

func (c *VOIPClient) NewAccountsAPI() *AccountsAPI {
	return &AccountsAPI{c}
}

func (c *VOIPClient) NewCDRAPI() *CDRAPI {
	return &CDRAPI{c}
}

func (c *VOIPClient) NewClientsAPI() *ClientsAPI {
	return &ClientsAPI{c}
}

//Only partially implemented. See dids.go.
func (c *VOIPClient) NewDIDsAPI() *DIDsAPI {
	return &DIDsAPI{c}
}

//Not implemented yet.
func (c *VOIPClient) NewFaxAPI() *FaxAPI {
	panic("NOT IMPLEMENTED YET!")
}

//Not implemented yet.
func (c *VOIPClient) NewVoicemailAPI() *VoicemailAPI {
	panic("NOT IMPLEMENTED YET!")
}

func (c *VOIPClient) WriteStruct(writer *multipart.Writer, iface interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(iface))

	FieldLoop:
	for i := 0; i < val.NumField(); i++ {

		structField := val.Type().Field(i)
		jsonTag := structField.Tag.Get("json")

		name := strings.TrimSuffix(jsonTag, ",omitempty")
		if name == "" {
			name = strings.ToLower(structField.Name)
		}

		value := ""
		omitEmpty := strings.Contains(jsonTag, ",omitempty")
		o := val.Field(i).Interface()

		t := structField.Type
		switch t.Kind() {
		case reflect.Struct:
			tm, ok := o.(encoding.TextMarshaler)
			if ok {
				text, err := tm.MarshalText()
				if err != nil {
					return err
				}

				value = string(text)
				break
			}

			//Write nested structs out and continue to the next field.
			//It is possible for nested struct field names to collide with top level field names with how this works.
			if err := c.WriteStruct(writer, o); err != nil {
				return err
			}
			continue FieldLoop
		case reflect.Bool: //Only write booleans if they are true. False is assumed if the field is not present in the request.
			b := o.(bool)
			if b {
				value = "true"
			} else {
				continue FieldLoop
			}
		case reflect.String: //Only write non-empty strings if they allow it with their JSON tags.
			value = fmt.Sprintf("%v", o) //json.Number and others also land here. Doing a o.(string) panics when that happens.
			if value == "" && omitEmpty {
				continue FieldLoop
			}
		default:
			value = fmt.Sprintf("%v", o)
			if c.Debug {
				log.Printf("Type: %s being written as value: %s for field: %s via WriteStruct:", t, value, name)
			}
		}

		if err := writer.WriteField(name, value); err != nil {
			return err
		}
	}

	return nil
}

