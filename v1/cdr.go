package v1

import (
	"net/url"
	"time"
	"errors"
	"fmt"
	"encoding/json"
	"strings"
)

type CDRAPI struct {
	client *VOIPClient
}

type CallAccount StringValueDescription

type GetCallAccountsResp struct {
	BaseResp
	CallAccounts []CallAccount `json:"accounts"`
}

type CallBilling StringValueDescription

type GetCallBillingResp struct {
	BaseResp
	CallBilling []CallBilling `json:"call_billing"`
}

type CallType StringValueDescription

type GetCallTypeResp struct {
	BaseResp
	CallTypes []CallType `json:"call_types"`
}

type CallStatus struct {
	Answered bool
	NoAnswer bool
	Busy     bool
	Failed   bool
}

type GetCDRResp struct {
	BaseResp
	CDRs []CDR `json:"cdr"`
}

type CDR struct {
	Date        time.Time `json:"date"`
	CallerId    string `json:"callerid"`
	Destination string `json:"destination"`
	Description string `json:"description"`
	Account     string `json:"account"`
	Disposition string `json:"disposition"`
	Duration    time.Duration `json:"duration"`
	Seconds     int `json:"seconds,string"`
	Rate        float64 `json:"rate,string,omitempty"`
	Total       float64 `json:"total,string"`
	UniqueId    string `json:"uniqueid"`
}

/*
func (c *CDR) MarshalJSON() ([]byte, error) {
	type Alias CDR
	
	str := c.Duration.String()
	str = strings.Replace(str, "h", ":", 1)
	str = strings.Replace(str, "m", ":", 1)
	str = strings.Replace(str, "s", ":", 1)
	
	return json.Marshal(&struct {
		Date string `json:"date"`
		Duration string `json:"duration"`
		*Alias
	}{
		Date: c.Date.Format("2006-01-02 15:04:05"),
		Duration: str,
		Alias:    (*Alias)(c),
	})
}
*/

func (c *CDR) UnmarshalJSON(data []byte) error {

	type Alias CDR
	aux := &struct {
		Date     string `json:"date"`
		Duration string `json:"duration"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	//date
	d, err := time.Parse("2006-01-02 15:04:05", aux.Date)
	if err != nil {
		return err
	}
	c.Date = d

	//duration
	str := strings.Split(aux.Duration, ":")
	fmt.Println(str)
	hour, err := time.ParseDuration(str[0] + "h")
	if err != nil {
		return err
	}

	min, err := time.ParseDuration(str[1] + "m")
	if err != nil {
		return err
	}

	sec, err := time.ParseDuration(str[2] + "s")
	if err != nil {
		return err
	}

	c.Duration = hour + min + sec

	return nil
}

type GetRatesResp struct {
	BaseResp
	Rates []Rate `json:"rates"`
}

type Rate struct {
	Destination     string `json:"destination"`
	Prefix          string `json:"prefix"`
	ClientIncrement int `json:"client_increment"`
	ClientRate      float64 `json:"client_rate"`
	RealIncrement   int `json:"real_increment"`
	RealRate        float64 `json:"real_rate"`
}

type GetTerminationRatesRep struct {
	BaseResp
	Route NumberValueDescription `json:"route"`
	Rates []TerminationRate `json:"rates"`
}

type TerminationRate struct {
	Destination string `json:"destination"`
	Prefix      string `json:"prefix"`
	Increment   int `json:"increment"`
	Rate        float64 `json:"rate"`
}

func (c *CDRAPI) GetCallAccounts(clientId string) ([]CallAccount, error) {
	values := url.Values{}
	if clientId != "" {
		values.Add("client", clientId)
	}

	rs := &GetCallAccountsResp{}
	if err := c.client.Get("getCallAccounts", values, rs); err != nil {
		return nil, err
	}

	return rs.CallAccounts, nil
}

func (c *CDRAPI) GetCallBilling() ([]CallBilling, error) {
	values := url.Values{}

	rs := &GetCallBillingResp{}
	if err := c.client.Get("getCallBilling", values, rs); err != nil {
		return nil, err
	}

	return rs.CallBilling, nil
}

func (c *CDRAPI) GetCallTypes(clientId string) ([]CallType, error) {
	values := url.Values{}
	if clientId != "" {
		values.Add("client", clientId)
	}

	rs := &GetCallTypeResp{}
	if err := c.client.Get("getCallTypes", values, rs); err != nil {
		return nil, err
	}

	return rs.CallTypes, nil
}

func (c *CDRAPI) GetCDR(dateFrom, dateTo time.Time, callStatus CallStatus, timezone *time.Location, callType, callBilling, account string) ([]CDR, error) {
	values, err := buildCDR(dateFrom, dateTo, callStatus, timezone, callType, callBilling, account)
	if err != nil {
		return nil, err
	}

	rs := &GetCDRResp{}
	if err := c.client.Get("getCDR", values, rs); err != nil {
		return nil, err
	}

	return rs.CDRs, nil
}

func (c *CDRAPI) GetRates(packag3, query string) ([]Rate, error) {
	values := url.Values{}
	values.Add("package", packag3)
	values.Add("query", query)

	rs := &GetRatesResp{}
	if err := c.client.Get("getRates", values, rs); err != nil {
		return nil, err
	}

	return rs.Rates, nil
}

func (c *CDRAPI) GetTerminationRates(route, query string) ([]TerminationRate, error) {
	values := url.Values{}
	values.Add("route", route)
	values.Add("query", query)

	rs := &GetTerminationRatesRep{}
	if err := c.client.Get("getTerminationRates", values, rs); err != nil {
		return nil, err
	}

	return rs.Rates, nil
}

func (c *CDRAPI) GetResellerCDR(dateFrom, dateTo time.Time, client string, callStatus CallStatus, timezone *time.Location, callType, callBilling, account string) ([]CDR, error) {
	values, err := buildCDR(dateFrom, dateTo, callStatus, timezone, callType, callBilling, account)
	if err != nil {
		return nil, err
	}

	values.Add("client", client)
	
	rs := &GetCDRResp{}
	if err := c.client.Get("getResellerCDR", values, rs); err != nil {
		return nil, err
	}

	return rs.CDRs, nil
}

func buildCDR(dateFrom, dateTo time.Time, callStatus CallStatus, timezone *time.Location, callType, callBilling, account string) (url.Values, error) {
	values := url.Values{}
	if dateFrom.IsZero() {
		return nil, errors.New("dateFrom is required!")
	}
	values.Add("date_from", dateFrom.Format("2006-01-02"))

	if dateTo.IsZero() {
		return nil, errors.New("dateTo is required!")
	}
	values.Add("date_to", dateTo.Format("2006-01-02"))

	if timezone.String() == "" {
		return nil, errors.New("timezone is required!")
	}
	_, offset := time.Now().In(timezone).Zone()
	d := time.Duration(offset) * time.Second
	values.Add("timezone", fmt.Sprintf("%.2g", d.Hours()))

	if callStatus.Answered {
		values.Add("answered", "1")
	}

	if callStatus.NoAnswer {
		values.Add("noanswer", "1")
	}

	if callStatus.Busy {
		values.Add("busy", "1")
	}

	if callStatus.Failed {
		values.Add("failed", "1")
	}

	values.Add("calltype", callType)
	values.Add("callbilling", callBilling)
	values.Add("account", account)

	return values, nil
}