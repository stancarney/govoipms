package govoipms

import (
	"fmt"
)

type GeneralAPI struct {
	client *Client
}

func NewGeneralAPI(client *Client) *GeneralAPI {
	return &GeneralAPI{client}
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

type Country StringValueDescription

type GetIPResp struct {
	BaseResp
	IP string `json:"ip"`
}

type GetLanguagesResp struct {
	BaseResp
	Languages []Language `json:"languages"`
}

type Language StringValueDescription

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

func (g *GeneralAPI) GetBalance(advanced bool) (*Balance, error) {
	url := g.client.BaseUrl("getBalance")

	if advanced {
		url = fmt.Sprintf("%s&advanced=true", url)
	}

	rs := &GetBalanceResp{}
	if err := g.client.Get(url, rs); err != nil {
		return nil, err
	}

	return &rs.Balance, nil
}

func (g *GeneralAPI) GetCountries(country string) ([]Country, error) {
	url := g.client.BaseUrl("getCountries")

	if country != "" {
		url = fmt.Sprintf("%s&country=%s", url, country)
	}

	rs := &GetCountriesResp{}
	if err := g.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Countries, nil
}

func (g *GeneralAPI) GetIP() (string, error) {
	url := g.client.BaseUrl("getIP")

	respStruct := &GetIPResp{}
	if err := g.client.Get(url, respStruct); err != nil {
		return "", err
	}

	return respStruct.IP, nil
}

func (g *GeneralAPI) GetLanguages(language string) ([]Language, error) {
	url := g.client.BaseUrl("getLanguages")

	if language != "" {
		url = fmt.Sprintf("%s&language=%s", url, language)
	}

	rs := &GetLanguagesResp{}
	if err := g.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Languages, nil
}

func (g *GeneralAPI) GetServerInfo(serverPop string) ([]Server, error) {
	url := g.client.BaseUrl("getServersInfo")

	if serverPop != "" {
		url = fmt.Sprintf("%s&server_pop=%s", url, serverPop)
	}

	rs := &GetServerInfoResp{}
	if err := g.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Servers, nil
}
