package govoipms

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
	"fmt"
	"errors"
)

type AccountAPI struct {
	client *Client
}

func NewAccountAPI(client *Client) *AccountAPI {
	return &AccountAPI{client}
}

type CreateSubAccountResp struct {
	BaseResp
	Output struct {
		       Id      string `json:"id"`
		       Account string `json:"account"`
	       }
}

type Account struct {
	Id                  string `json:"id,omitempty"`
	Account             string `json:"account,omitempty"`
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

type GetAllowedCodecsResp struct {
	BaseResp
	AllowedCodecs []Codec `json:"allowed_codecs"`
}

type Codec StringValueDescription

type GetAuthTypesResp struct {
	BaseResp
	AuthTypes []AuthType `json:"auth_types"`
}

type AuthType NumberValueDescription

type GetDeviceTypesResp struct {
	BaseResp
	DeviceTypes []DeviceType `json:"device_types"`
}

type DeviceType NumberValueDescription

type GetDTMFModesResp struct {
	BaseResp
	DTMFModes []DTMFMode `json:"dtmf_modes"`
}

type DTMFMode StringValueDescription

type GetLockInternationalResp struct {
	BaseResp
	LockInternational []LockInternational `json:"lock_international"`
}

//The Value type seems to switch between int and string for some reason, so it needs a custom struct for now. TODO:Re-evaluate
type LockInternational NumberValueDescription

type GetMusicOnHoldResp struct {
	BaseResp
	MusicOnHold []MusicOnHold `json:"music_on_hold"`
}

type MusicOnHold StringValueDescription

type GetNATResp struct {
	BaseResp
	NAT []NAT `json:"nat"`
}

type NAT StringValueDescription

type GetProtocolResp struct {
	BaseResp
	Protocols []Protocol `json:"protocols"`
}

type Protocol NumberValueDescription

type GetRegistrationStatusResp struct {
	BaseResp
	Registered    string `json:"registered"`
	Registrations []RegistrationStatus `json:"registrations"`
}

type RegistrationStatus struct {
	Server
	RegisterIP   string `json:"register_ip"`
	RegisterPort string `json:"register_port"`
	RegisterNext string `json:"register_next"`
}

type GetReportEstimatedHoldTimeResp struct {
	BaseResp
	Types []EstimatedHoldTime `json:"types"`
}

type EstimatedHoldTime StringValueDescription

type GetRoutesResp struct {
	BaseResp
	Routes []Route `json:"routes"`
}

type Route NumberValueDescription

type GetSubAccountsResp struct {
	BaseResp
	Accounts []Account `json:"accounts"`
}

type SetSubAccountsResp BaseResp

func (a *AccountAPI) CreateSubAccount(subAccount *Account) (*Balance, error) {
	url := a.client.BaseUrl("createSubAccount")

	b, err := json.Marshal(subAccount)

	if a.client.debug {
		log.Println(string(b))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	respStruct := &GetBalanceResp{}
	_, err = a.client.Call(req, respStruct)
	if err != nil {
		return nil, err
	}

	if respStruct.Status != "success" {
		return nil, errors.New("Not Succesful!")
	}

	return &respStruct.Balance, nil
}

func (a *AccountAPI) DelSubAccount(subAccount *Account) *Balance {
	panic("NOT IMPLEMENTED YET!")
}

func (a *AccountAPI) GetAllowedCodecs(codec string) ([]Codec, error) {
	url := a.client.BaseUrl("getAllowedCodecs")

	if codec != "" {
		url = fmt.Sprintf("%s&codec=%s", url, codec)
	}

	rs := &GetAllowedCodecsResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.AllowedCodecs, nil
}

func (a *AccountAPI) GetAuthTypes(authType int) ([]AuthType, error) {
	url := a.client.BaseUrl("getAuthTypes")

	if authType > 0 {
		url = fmt.Sprintf("%s&type=%d", url, authType)
	}

	rs := &GetAuthTypesResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.AuthTypes, nil
}

func (a *AccountAPI) GetDeviceTypes(deviceType int) ([]DeviceType, error) {
	url := a.client.BaseUrl("getDeviceTypes")

	if deviceType > 0 {
		url = fmt.Sprintf("%s&device_type=%d", url, deviceType)
	}

	rs := &GetDeviceTypesResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.DeviceTypes, nil
}

func (a *AccountAPI) GetDTMFModes(DTMFMode string) ([]DTMFMode, error) {
	url := a.client.BaseUrl("getDTMFModes")

	if DTMFMode != "" {
		url = fmt.Sprintf("%s&dtmf_mode=%s", url, DTMFMode)
	}

	rs := &GetDTMFModesResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.DTMFModes, nil
}

//0 is an actual value for a Lock International entity so the signature of this message is a string opposed to an int.
//This was done to avoid confusion with other functions that take 0 in order to return all values.
func (a *AccountAPI) GetLockInternational(lockInternational string) ([]LockInternational, error) {
	url := a.client.BaseUrl("getLockInternational")

	if lockInternational != "" {
		url = fmt.Sprintf("%s&lock_international=%s", url, lockInternational)
	}

	rs := &GetLockInternationalResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.LockInternational, nil
}

func (a *AccountAPI) GetMusicOnHold(musicOnHold string) ([]MusicOnHold, error) {
	url := a.client.BaseUrl("getMusicOnHold")

	if musicOnHold != "" {
		url = fmt.Sprintf("%s&music_on_hold=%s", url, musicOnHold)
	}

	rs := &GetMusicOnHoldResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.MusicOnHold, nil
}

func (a *AccountAPI) GetNAT(NAT string) ([]NAT, error) {
	url := a.client.BaseUrl("getNAT")

	if NAT != "" {
		url = fmt.Sprintf("%s&nat=%s", url, NAT)
	}

	rs := &GetNATResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.NAT, nil
}

func (a *AccountAPI) GetProtocols(protocol int) ([]Protocol, error) {
	url := a.client.BaseUrl("getProtocols")

	if protocol > 0 {
		url = fmt.Sprintf("%s&protocol=%d", url, protocol)
	}

	rs := &GetProtocolResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Protocols, nil
}

func (a *AccountAPI) GetRegistrationStatus(account string) (bool, []RegistrationStatus, error) {
	url := a.client.BaseUrl("getRegistrationStatus")

	if account == "" {
		return false, nil, errors.New("missing_account")
	}

	url = fmt.Sprintf("%s&account=%s", url, account)

	rs := &GetRegistrationStatusResp{}
	if err := a.client.Get(url, rs); err != nil {
		return false, nil, err
	}

	return rs.Registered == "yes", rs.Registrations, nil
}

func (a *AccountAPI) GetReportEstimatedHoldTime(typ3 string) ([]EstimatedHoldTime, error) {
	url := a.client.BaseUrl("getReportEstimatedHoldTime")

	if typ3 != "" {
		url = fmt.Sprintf("%s&type=%s", url, typ3)
	}

	rs := &GetReportEstimatedHoldTimeResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Types, nil
}

func (a *AccountAPI) GetRoutes(route int) ([]Route, error) {
	url := a.client.BaseUrl("getRoutes")

	if route > 0 {
		url = fmt.Sprintf("%s&route=%d", url, route)
	}

	rs := &GetRoutesResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Routes, nil
}

func (a *AccountAPI) GetSubAccounts(account string) ([]Account, error) {
	url := a.client.BaseUrl("getSubAccounts")

	if account != "" {
		url = fmt.Sprintf("%s&account=%s", url, account)
	}

	rs := &GetSubAccountsResp{}
	if err := a.client.Get(url, rs); err != nil {
		return nil, err
	}

	return rs.Accounts, nil
}

func (a *AccountAPI) SetSubAccount(account *Account) error {
	url := a.client.BaseUrl("setSubAccount")

	rs := &SetSubAccountsResp{}
	if err := a.client.Get(url, rs); err != nil {
		return err
	}

	return nil
}