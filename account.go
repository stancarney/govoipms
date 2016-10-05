package govoipms

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

type Account struct {
	client *Client
}

func NewAccount(client *Client) *Account {
	return &Account{client}
}

type CreateSubAccountResp struct {
	BaseResp
	Output struct {
		       Id      string `json:"id"`
		       Account string `json:"account"`
	       }
}

type SubAccount struct {
	Username            string `json:"username"`
	Protocol            string `json:"protocol"`
	Description         string `json:"description"`
	AuthType            string `json:"auth_type"`
	Password            string `json:"password"`
	IP                  string `json:"ip"`
	DeviceType          string `json:"device_type"`
	CalleridNumber      string `json:"callerid_number"`
	CanadaRouting       string `json:"canada_routing"`
	LockInternational   string `json:"lock_international"`
	InternationalRoute  string `json:"international_route"`
	MusicOnHold         string `json:"music_on_hold"`
	AllowedCodecs       string `json:"allowed_codecs"`
	DTMFMode            string `json:"dtmf_mode"`
	NAT                 string `json:"nat"`
	InternalExtension   string `json:"internal_extension"`
	InternalVoicemail   string `json:"internal_voicemail"`
	InternalDialtime    string `json:"internal_dialtime"`
	ResellerClient      string `json:"reseller_client"`
	ResellerPackage     string `json:"reseller_package"`
	ResellerNextBilling string `json:"reseller_nextbilling"`
	ResellerChargesetup string `json:"reseller_chargesetup"`
}

func (g *General) CreateSubAccount(subAccount *SubAccount) *Balance {
	url := g.client.BaseUrl("createSubAccount")

	b, err := json.Marshal(subAccount)

	if g.client.debug {
		log.Println(string(b))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
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
