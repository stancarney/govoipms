package v1

import (
	"net/url"
	"encoding/json"
	"strings"
)

type DIDsAPI struct {
	client *VOIPClient
}

type BaseRoute struct {
	Type  string
	Value string
}

func (b BaseRoute) String() string {
	return strings.Join([]string{b.Type, b.Value}, ":")
}

func (b BaseRoute) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

//Helper to create account route.
func NewAccountRoute(value string) BaseRoute {
	return BaseRoute{"account", value}
}

//Helper to create fwd route.
func NewFwdRoute(value string) BaseRoute {
	return BaseRoute{"fwd", value}
}

//Helper to create vm route.
func NewVMRoute(value string) BaseRoute {
	return BaseRoute{"vm", value}
}

//Helper to create sys route.
func NewSysRoute(value string) BaseRoute {
	return BaseRoute{"sys", value}
}

//Helper to create none route.
func NewNoneRoute() BaseRoute {
	return BaseRoute{"none", ""}
}

type BackOrder struct {
	Quantity            int `json:"quantity"`
	State               string `json:"state,omitempty"`    //State is used for US functions.
	Province            string `json:"province,omitempty"` //Province is used for CA functions.
	Ratecenter          string `json:"ratecenter"`
	Routing             BaseRoute `json:"routing"`
	FailoverBusy        BaseRoute `json:"failover_busy,omitempty"`
	FailoverUnreachable BaseRoute `json:"failover_unreachable,omitempty"`
	FailoverNoanswer    BaseRoute `json:"failover_noanswer,omitempty"`
	Voicemail           string `json:"voicemail,omitempty"`
	POP                 string `json:"pop"`
	Dialtime            int `json:"dialtime"`
	CNAM                bool `json:"cnam"`
	CalleridPrefix      string `json:"callerid_prefix,omitempty"`
	Note                string `json:"note,omitempty"`
	BillingType         int `json:"billing_type"`
	Test                bool `json:"test"`
}

type GetRateCentersResp struct {
	BaseResp
	RateCenters []RateCenter `json:"ratecenters"`
}

type RateCenter struct {
	RateCenter string `json:"ratecenter"`
	Available  bool `json:"available"`
}

type GetStatesResp struct {
	BaseResp
	States []State `json:"states"`
}

type State struct {
	State       string `json:"state"`
	Description string `json:"description"`
}

func (c *RateCenter) UnmarshalJSON(data []byte) error {

	type Alias RateCenter
	aux := &struct {
		Available string `json:"available"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Available == "yes" {
		c.Available = true
	}

	return nil
}

//TODO:Stan this isn't working. It returns "invalid_ratecenter" and I'm pretty sure the ratecenter is correct.
func (d *DIDsAPI) BackOrderDIDUSA(backOrder *BackOrder) error {
	rs := &BaseResp{}
	rq := backOrder

	if err := d.client.Post("backOrderDIDUSA", rq, rs); err != nil {
		return err
	}

	return nil
}

//TODO:Stan this isn't working. It returns "invalid_ratecenter" and I'm pretty sure the ratecenter is correct.
func (d *DIDsAPI) BackOrderDIDCAN(backOrder *BackOrder) error {
	rs := &BaseResp{}
	rq := backOrder

	if err := d.client.Post("backOrderDIDCAN", rq, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) GetRateCentersCAN(province string) ([]RateCenter, error) {
	values := url.Values{}
	values.Add("province", province)

	rs := &GetRateCentersResp{}
	if err := d.client.Get("getRateCentersCAN", values, rs); err != nil {
		return nil, err
	}

	return rs.RateCenters, nil
}

func (d *DIDsAPI) GetRateCentersUSA(state string) ([]RateCenter, error) {
	values := url.Values{}
	values.Add("state", state)

	rs := &GetRateCentersResp{}
	if err := d.client.Get("getRateCentersUSA", values, rs); err != nil {
		return nil, err
	}

	return rs.RateCenters, nil
}

func (d *DIDsAPI) GetStates() ([]State, error) {
	rs := &GetStatesResp{}
	if err := d.client.Get("getStates", url.Values{}, rs); err != nil {
		return nil, err
	}

	return rs.States, nil
}