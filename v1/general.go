package v1

import (
	"net/url"
)

type GeneralAPI struct {
	client *VOIPClient
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

	values := url.Values{}
	if advanced {
		values.Add("advanced", "true")
	}

	rs := &GetBalanceResp{}
	if err := g.client.Get("getBalance", values, rs); err != nil {
		return nil, err
	}

	return &rs.Balance, nil
}

func (g *GeneralAPI) GetCountries(country string) ([]Country, error) {
	values := url.Values{}
	if country != "" {
		values.Add("country", country)
	}

	rs := &GetCountriesResp{}
	if err := g.client.Get("getCountries", values, rs); err != nil {
		return nil, err
	}

	return rs.Countries, nil
}

func (g *GeneralAPI) GetIP() (string, error) {
	respStruct := &GetIPResp{}
	if err := g.client.Get("getIP", url.Values{}, respStruct); err != nil {
		return "", err
	}

	return respStruct.IP, nil
}

func (g *GeneralAPI) GetLanguages(language string) ([]Language, error) {
	values := url.Values{}
	if language != "" {
		values.Add("language", language)
	}

	rs := &GetLanguagesResp{}
	if err := g.client.Get("getLanguages", values, rs); err != nil {
		return nil, err
	}

	return rs.Languages, nil
}

func (g *GeneralAPI) GetServerInfo(serverPop string) ([]Server, error) {
	values := url.Values{}
	if serverPop != "" {
		values.Add("server_pop", serverPop)
	}

	rs := &GetServerInfoResp{}
	if err := g.client.Get("getServersInfo", values, rs); err != nil {
		return nil, err
	}

	return rs.Servers, nil
}
