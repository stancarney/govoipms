package v1

import (
	"net/url"
	"time"
	"errors"
	"encoding/json"
)

type GeneralAPI struct {
	client *VOIPClient
}

type GetBalanceResp struct {
	BaseResp
	Balance `json:"balance"`
}

type Balance struct {
	CurrentBalance json.Number `json:"current_balance"`
	SpentTotal     json.Number `json:"spent_total,omitempty"`
	CallsTotal     json.Number `json:"calls_total,omitempty"`
	TimeTotal      json.Number `json:"time_total,omitempty"` //TODO:Stan change to duration
	SpentToday     json.Number `json:"spent_today,omitempty"`
	CallsToday     int `json:"calls_today,omitempty"`
	TimeToday      json.Number `json:"time_today,omitempty"`
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

type GetTransactionHistoryResp struct {
	BaseResp
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Date        string `json:"date"` //can't be time.Time as the API likes to return things like this: {"date":"2015-11-02 to 2016-11-01","uniqueid":"n\/a","type":"CNAM Queries","description":"CNAM Queries","ammount":"-0.2160"}
	UniqueId    string `json:"uniqueid"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
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

func (g *GeneralAPI) GetTransactionHistory(dateFrom, dateTo time.Time) ([]Transaction, error) {
	values := url.Values{}
	if dateFrom.IsZero() {
		return nil, errors.New("dateFrom is required!")
	}
	values.Add("date_from", dateFrom.Format("2006-01-02 15:04:05"))

	if dateTo.IsZero() {
		return nil, errors.New("dateTo is required!")
	}
	values.Add("date_to", dateTo.Format("2006-01-02 15:04:05"))

	rs := &GetTransactionHistoryResp{}
	if err := g.client.Get("getTransactionHistory", values, rs); err != nil {
		return nil, err
	}

	return rs.Transactions, nil
}
