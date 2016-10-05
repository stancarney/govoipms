package vclient

import (
	"net/http"
	"fmt"
)

type General struct {
	client *Client
}

func NewGeneral(client *Client) *General {
	return &General{client}
}

type GetBalanceResp struct {
	BaseResp
	Balance `json:"balance"`
}

type Balance struct {
	CurrentBalance string `json:"current_balance"`
	SpentTotal     float64 `json:"spent_total,omitempty"`
	CallsTotal     float64 `json:"calls_total,omitempty"`
	TimeTotal      string `json:"time_total,omitempty"`
	SpentToday     float64 `json:"spent_today,omitempty"`
	CallsToday     int `json:"calls_today,omitempty"`
	TimeToday      string `json:"time_today,omitempty"`
}

type GetCountriesResp struct {
	BaseResp
	Countries []Country `json:"countries"`
}

type Country ValueDescription

type GetIPResp struct {
	BaseResp
	IP string `json:"ip"`
}

type GetLanguagesResp struct {
	BaseResp
	Languages []Language `json:"languages"`
}

type Language ValueDescription

type GetServerInfoResp struct {
	BaseResp
	Servers []Server `json:"servers"`
}

type Server struct {
	ServerName      string `json:"server_name"`
	ServerShortname string `json:"server_shortname"`
	ServerHostname  string `json:"server_hostname"`
	ServerIP        string `json:"server_ip"`
	ServerCountry   string `json:"server_country"`
	ServerPop       string `json:"server_pop"`
}

func (g *General) GetBalance(advanced bool) *Balance {
	url := g.client.BaseUrl("getBalance")

	if advanced {
		url = fmt.Sprintf("%s&advanced=true", url)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	respStruct := &GetBalanceResp{}
	_, err = g.client.Call(req, respStruct)
	if err != nil {
		panic(err)
	}

	if respStruct.Status != "success" {
		panic("Not successful!")
	}

	return &respStruct.Balance
}

func (g *General) GetCountries(country string) []Country {
	url := g.client.BaseUrl("getCountries")

	if country != "" {
		url = fmt.Sprintf("%s&country=%s", url, country)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	respStruct := &GetCountriesResp{}
	_, err = g.client.Call(req, respStruct)
	if err != nil {
		panic(err)
	}

	if respStruct.Status != "success" {
		panic("Not successful!")
	}

	return respStruct.Countries
}

func (g *General) GetIP() string {
	url := g.client.BaseUrl("getIP")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	respStruct := &GetIPResp{}
	_, err = g.client.Call(req, respStruct)
	if err != nil {
		panic(err)
	}

	if respStruct.Status != "success" {
		panic("Not successful!")
	}

	return respStruct.IP
}

func (g *General) GetLanguages(language string) []Language {
	url := g.client.BaseUrl("getLanguages")

	if language != "" {
		url = fmt.Sprintf("%s&language=%s", url, language)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	respStruct := &GetLanguagesResp{}
	_, err = g.client.Call(req, respStruct)
	if err != nil {
		panic(err)
	}

	if respStruct.Status != "success" {
		panic("Not successful!")
	}

	return respStruct.Languages
}

func (g *General) GetServerInfo(serverPop string) []Server {
	url := g.client.BaseUrl("getServersInfo")

	if serverPop != "" {
		url = fmt.Sprintf("%s&server_pop=%s", url, serverPop)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	respStruct := &GetServerInfoResp{}
	_, err = g.client.Call(req, respStruct)
	if err != nil {
		panic(err)
	}

	if respStruct.Status != "success" {
		panic("Not successful!")
	}

	return respStruct.Servers
}
