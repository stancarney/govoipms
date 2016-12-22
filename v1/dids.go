package v1

import (
	"net/url"
	"encoding/json"
	"strings"
	"errors"
	"time"
	"fmt"
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

func (b *BaseRoute) UnmarshalText(text []byte) error {
	strs := strings.Split(string(text), ":")

	l := len(strs)

	if l > 0 {
		b.Type = strs[0]
	}

	if l > 1 {
		b.Value = strs[1]
	}

	return nil
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

//Helper to create sip route.
func NewSIPRoute() BaseRoute {
	return BaseRoute{"sip", ""}
}

//Helper to create grp route.
func NewGrpRoute() BaseRoute {
	return BaseRoute{"grp", ""}
}

//Helper to create ivr route.
func NewIVRRoute() BaseRoute {
	return BaseRoute{"ivr", ""}
}

//Helper to create sys route.
func NewSysRoute(value string) BaseRoute {
	return BaseRoute{"sys", value}
}

//Helper to create recording route.
func NewRecordingRoute() BaseRoute {
	return BaseRoute{"recording", ""}
}

//Helper to create queue route.
func NewQueueRoute() BaseRoute {
	return BaseRoute{"queue", ""}
}

//Helper to create cb route.
func NewCBRoute() BaseRoute {
	return BaseRoute{"cb", ""}
}

//Helper to create tc route.
func NewTCRoute() BaseRoute {
	return BaseRoute{"tc", ""}
}

//Helper to create disa route.
func NewDisaRoute() BaseRoute {
	return BaseRoute{"disa", ""}
}

//Helper to create none route.
func NewNoneRoute() BaseRoute {
	return BaseRoute{"none", ""}
}

type BackOrder struct {
	Quantity   int `json:"quantity"`
	State      string `json:"state,omitempty"`    //State is used for US functions.
	Province   string `json:"province,omitempty"` //Province is used for CA functions.
	Ratecenter string `json:"ratecenter"`
	Order
}

type Order struct {
	Routing             BaseRoute `json:"routing"`
	FailoverBusy        BaseRoute `json:"failover_busy,omitempty"`
	FailoverUnreachable BaseRoute `json:"failover_unreachable,omitempty"`
	FailoverNoanswer    BaseRoute `json:"failover_noanswer,omitempty"`
	Voicemail           string `json:"voicemail,omitempty"`
	POP                 string `json:"pop"`
	Dialtime            json.Number `json:"dialtime"`
	CNAM                json.Number `json:"cnam"` //1: true, 0: false.
	CalleridPrefix      string `json:"callerid_prefix,omitempty"`
	Note                string `json:"note,omitempty"`
	BillingType         json.Number `json:"billing_type"`
	Test                bool `json:"test,omitempty"`
}

type CancelDIDResp struct {
	BaseResp
}

type ConnectDIDResp struct {
	BaseResp
}

type GetDIDCountriesResp struct {
	BaseResp
	Countries []DIDCountries `json:"countries"`
}

type DIDCountries NumberValueDescription

type GetDIDsCanResp struct {
	BaseResp
	DIDs []DID `json:"dids"`
}

type GetDIDsUSAResp struct {
	BaseResp
	DIDs []DID `json:"dids"`
}

type GetDISAsResp struct {
	BaseResp
	DISAs []DISA `json:"disa"`
}

type DISA struct {
	DISA             string `json:"disa"`
	Name             string `json:"name"`
	PIN              string `json:"pin"`
	DigitTimeout     string `json:"digit_timeout"`
	CalleridOverride string `json:"callerid_override"`
}

type GetForwardingsResp struct {
	BaseResp
	Forwardings []Forwarding `json:"forwardings"`
}

type Forwarding struct {
	Forwarding       string `json:"forwarding"`
	PhoneNumber      string `json:"phone_number"`
	CalleridOverride string `json:"callerid_override"`
	Description      string `json:"description"`
	DTMFDigits       string `json:"dtmf_digits"`
	Pause            string `json:"pause"`
}

type GetIVRsResp struct {
	BaseResp
	IVRs []IVR `json:"ivrs"`
}

type IVR struct {
	IVR            string `json:"ivr"`
	Name           string `json:"name"`
	Recording      string `json:"recording"`
	Timeout        string `json:"timeout"`
	Language       string `json:"language"`
	VoicemailSetup string `json:"voicemailsetup"`
	Choices        []IVRChoice `json:"choices"`
}

func (i *IVR) UnmarshalJSON(data []byte) error {
	type Alias IVR
	aux := &struct {
		Choices string `json:"choices"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	//This converts the incoming string "1=account:12345;2=queue:3333" to IVRChoice Records
	strs := strings.Split(aux.Choices, ";")
	i.Choices = make([]IVRChoice, len(strs))

	for idx, c := range strs {
		i.Choices[idx] = IVRChoice{}
		if err := i.Choices[idx].UnmarshalText([]byte(c)); err != nil {
			return err
		}
	}

	return nil
}

type IVRChoice struct {
	DTMFTone string
	Route    BaseRoute
}

func (i IVRChoice) String() string {
	return fmt.Sprintf("%s=%s", i.DTMFTone, i.Route)
}

func (i IVRChoice) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *IVRChoice) UnmarshalText(text []byte) error {
	strs := strings.Split(string(text), "=")

	l := len(strs)

	if l > 0 {
		i.DTMFTone = strs[0]
	}

	if l > 1 {
		i.Route = NewNoneRoute()
		if err := i.Route.UnmarshalText([]byte(strs[1])); err != nil {
			return err
		}
	}

	return nil
}

type GetJoinWhenEmptyTypesResp struct {
	BaseResp
	JoinWhenEmptyTypes []JoinWhenEmptyType `json:"types"`
}

type JoinWhenEmptyType StringValueDescription

type GetPhonebookResp struct {
	BaseResp
	Phonebooks []Phonebook `json:"phonebooks"`
}

type Phonebook struct {
	Phonebook string `json:"phonebook"`
	SpeedDial string `json:"speed_dial"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Callerid  string `json:"callerid"`
	Note      string `json:"note"`
}

type GetPortabilityResp struct {
	BaseResp
	Portable string `json:"portable"`
	Plans    []Plan `json:"plans"`
}

type Plan struct {
	Title         string `json:"title"`
	PricePerMonth string `json:"pricePerMonth"` //Voip.MS API uses camelcase here instead of underscores for some reason.
	PricePerMin   string `json:"pricePerMin"`
}

type GetProvincesResp struct {
	BaseResp
	Provinces []Province `json:"provinces"`
}

type Province struct {
	Province    string `json:"province"`
	Description string `json:"description"`
}

type GetDIDsInfoResp struct {
	BaseResp
	DIDInfos []DIDInfo `json:"dids"`
}

type DIDInfo struct {
	DID                   string `json:"did"`
	Description           string `json:"description"`
	Order
	E911                  string `json:"e911"`
	NextBilling           time.Time `json:"next_billing"`
	OrderDate             time.Time `json:"next_billing"`
	ResellerAccount       string `json:"reseller_account"`

	ResellerNextBilling   string `json:"reseller_next_billing"`
	ResellerMonthly       string `json:"reseller_monthly"`
	ResellerMinute        string `json:"reseller_minute"`
	ResellerSetup         string `json:"reseller_setup"`

	SMSAvailable          int `json:"sms_available"`
	SMSEnabled            string `json:"sms_enabled"`
	SMSEmail              string `json:"sms_email"`
	SMSEmailEnabled       string `json:"sms_email_enabled"`
	SMSForward            string `json:"sms_forward"`
	SMSForwardEnabled     string `json:"sms_forward_enabled"`
	SMSURLCallback        string `json:"sms_url_callback"`
	SMSURLCallbackEnabled string `json:"sms_url_callback_enabled"`
	SMSURLCallbackEntry   string `json:"sms_url_callback_retry"`
}

type DelStaticMemberResp struct {
	BaseResp
}

type GetCallbacksResp struct {
	BaseResp
	Callbacks []Callback `json:"callbacks"`
}

type Callback struct {
	Callback        string `json:"callback"`
	Description     string `json:"description"`
	Number          string `json:"number"`
	DelayBefore     int `json:"delay_before"`
	ResponseTimeout int `json:"response_timeout"`
	DigitTimeout    int `json:"digit_timeout"`
	CalleridNumber  string `json:"callerid_number"`
}

type GetCallerIDFilteringResp struct {
	BaseResp
	CallerIDFilters []CallerIDFilter `json:"filtering"`
}

type CallerIDFilter struct {
	Filtering           string `json:"filtering"`
	Callerid            string `json:"callerid"`
	DID                 string `json:"did"`
	Routing             string `json:"routing"`
	FailoverUnreachable string `json:"failover_unreachable"`
	FailoverBusy        string `json:"failover_busy"`
	FailoverNoanswer    string `json:"failover_noanswer"`
	Note                string `json:"note"`
}

type GetCarriersResp struct {
	BaseResp
	Carriers []Carrier `json:"carriers"`
}

type Carrier NumberValueDescription

type GetDIDsInternationalResp struct {
	BaseResp
	Locations []InternationalLocations `json:"locations"`
}

type InternationalLocations struct {
	LocationId   string `json:"location_id"`
	LocationName string `json:"location_name"`
	Country      string `json:"country"`
	AreaCode     string `json:"area_code"`
	Stock        string `json:"stock"`
	Monthly      string `json:"montly"`
	Setup        string `json:"setup"`
	Minute       string `json:"minute"`
	Channels     string `json:"channels,omitempty"` //only used for GetDIDsInternationalGeographic
}

type GetInternationalTypesResp struct {
	BaseResp
	Types []InternationalTypes `json:"types"`
}

type InternationalTypes StringValueDescription

type GetRateCentersResp struct {
	BaseResp
	RateCenters []RateCenter `json:"ratecenters"`
}

type RateCenter struct {
	RateCenter string `json:"ratecenter"`
	Available  bool `json:"available"`
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

type GetStatesResp struct {
	BaseResp
	States []State `json:"states"`
}

type State struct {
	State       string `json:"state"`
	Description string `json:"description"`
}

type GetStaticMembersResp struct {
	BaseResp
	Members []Member `json:"members"`
}

type Member struct {
	Member    string `json:"member"`
	QueueName string `json:"queue_name"`
	Name      string `json:"name"`
	Account   string `json:"account"`
	Priority  string `json:"priority"`
}

type GetTimeConditionsResp struct {
	BaseResp
	TimeConditions []TimeCondition `json:"timeconditon"`
}

type TimeCondition struct {
	TimeCondition  string `json:"timecondition"`
	Name           string `json:"name"`
	RoutingMatch   string `json:"routingmatch"`
	RoutingNoMatch string `json:"routingnomatch"`
	StartHour      string `json:"starthour"`
	StartMinute    string `json:"startminute"`
	EndHour        string `json:"endhour"`
	EndMinute      string `json:"endminute"`
	WeekdayStart   string `json:"weekdaystart"`
	WeekdayEnd     string `json:"weekdayend"`
}

type GetVoicemailSetups struct {
	BaseResp
	VoicemailSetups []VoicemailSetup `json:"voicemailsetups"`
}

type VoicemailSetup NumberValueDescription

type GetVoicemailAttachmentFormats struct {
	BaseResp
	VoicemailAttachmentFormats []VoicemailAttachmentFormat `json:"email_attachment_formats"`
}

type VoicemailAttachmentFormat StringValueDescription

type DIDOrderResellerConfig struct {
	Account string `json:"account"`
	Monthly string `json:"monthly"`
	Setup   string `json:"setup"`
	Minute  string `json:"minute"`
}

type DIDOrder struct {
	Did string `json:"did"`
	Order
	DIDOrderResellerConfig
}

type DIDOrderInternationalGeographic struct {
	LocationId string `json:"location_id"`
	Quantity   int `json:"quantity"`
	Order
	DIDOrderResellerConfig
}

type SearchDIDsCanResp struct {
	BaseResp
	DIDs []DID `json:"dids"`
}

type DID struct {
	DID                 string `json:"did"`
	Ratecenter          string `json:"ratecenter"`
	Province            string `json:"province,omitempty"`             //only populated on Canadian API calls
	ProvinceDescription string `json:"province_description,omitempty"` //only populated on Canadian API calls
	State               string `json:"state,omitempty"`                //only populated on US API calls
	StateDescription    string `json:"state_description,omitempty"`    //only populated on US API calls
	PerMinuteMonthly    string `json:"perminute_monthly"`
	PerMinuteMinute     string `json:"perminute_minute"`
	PerMinuteSetup      string `json:"perminute_setup"`
	FlatMonthly         string `json:"flat_monthly"`
	FlatMinute          string `json:"flat_minute"`
	FlatSetup           string `json:"flat_setup"`
	SMS                 int `json:"sms,omitempty"`
}

type DIDSearchType string

const (
	StartsDIDSearchType DIDSearchType = "starts"
	ContainsDIDSearchType DIDSearchType = "contains"
	EndsDIDSearchType DIDSearchType = "ends"
)

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
func (d *DIDsAPI) BackOrderDIDCan(backOrder *BackOrder) error {
	rs := &BaseResp{}
	rq := backOrder

	if err := d.client.Post("backOrderDIDCAN", rq, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) CancelDID(DID, comment string, portOut, test bool) error {
	values := url.Values{}
	values.Add("did", DID)

	if portOut {
		values.Add("portout", "true")
	}

	if test {
		values.Add("test", "true")
	}

	if comment != "" {
		values.Add("cancelcomment", comment)
	}

	rs := &CancelDIDResp{}
	//TODO:Stan this is called "CancelDID" in the documentation...
	if err := d.client.Get("cancelDID", values, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) ConnectDID(DID, account, monthly, setup, minute string, nextBilling time.Time, dontChargeSetup, dontChargeMonthly bool) error {
	values := url.Values{}
	values.Add("did", DID)
	values.Add("account", account)
	values.Add("monthly", monthly)
	values.Add("setup", setup)
	values.Add("minute", minute)

	if !nextBilling.IsZero() {
		values.Add("next_billing", nextBilling.Format("2006-01-02"))
	}

	if dontChargeSetup {
		values.Add("dont_charge_setup", "true")
	}

	if dontChargeMonthly {
		values.Add("dont_charge_monthly", "true")
	}

	rs := &ConnectDIDResp{}
	if err := d.client.Get("connectDID", values, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) DelCallback(callback string) error {
	return d.client.simpleCall("delCallback", "callback", callback)
}

func (d *DIDsAPI) DelCallerIDFiltering(filtering string) error {
	return d.client.simpleCall("delCallerIDFiltering", "filtering", filtering)
}

func (d *DIDsAPI) DelClient(client string) error {
	return d.client.simpleCall("delClient", "client", client)
}

func (d *DIDsAPI) DelDISA(disa string) error {
	return d.client.simpleCall("delDISA", "disa", disa)
}

func (d *DIDsAPI) DeleteSMS(id string) error {
	return d.client.simpleCall("deleteSMS", "id", id)
}

func (d *DIDsAPI) DelForwarding(forwarding string) error {
	return d.client.simpleCall("delForwarding", "forwarding", forwarding)
}

func (d *DIDsAPI) DelIVR(ivr string) error {
	return d.client.simpleCall("delIVR", "ivr", ivr)
}

func (d *DIDsAPI) DelPhonebook(phonebook string) error {
	return d.client.simpleCall("delPhonebook", "phonebook", phonebook)
}

func (d *DIDsAPI) DelQueue(queue string) error {
	return d.client.simpleCall("delQueue", "queue", queue)
}

func (d *DIDsAPI) DelRecording(recording string) error {
	return d.client.simpleCall("delRecording", "recording", recording)
}

func (d *DIDsAPI) DelRingGroup(ringGroup string) error {
	return d.client.simpleCall("delRingGroup", "ringGroup", ringGroup)
}

func (d *DIDsAPI) DelSIPURI(SIPURI string) error {
	return d.client.simpleCall("delSIPURI", "sipuri", SIPURI)
}

func (d *DIDsAPI) DelStaticMember(member, queue string) error {
	values := url.Values{}
	values.Add("member", member)
	values.Add("queue", queue)

	rs := &DelStaticMemberResp{}
	return d.client.Get("delStaticMember", values, rs)
}

func (d *DIDsAPI) DelTimeCondition(timeCondition string) error {
	return d.client.simpleCall("delTimeCondition", "timecondition", timeCondition)
}

func (d *DIDsAPI) GetCallbacks(callback string) ([]Callback, error) {
	values := url.Values{}

	if callback != "" {
		values.Add("callback", callback)
	}

	rs := &GetCallbacksResp{}
	if err := d.client.Get("getCallbacks", values, rs); err != nil {
		return nil, err
	}

	return rs.Callbacks, nil
}

func (d *DIDsAPI) GetCallerIDFiltering(filtering string) ([]CallerIDFilter, error) {
	values := url.Values{}

	if filtering != "" {
		values.Add("filtering", filtering)
	}

	rs := &GetCallerIDFilteringResp{}
	if err := d.client.Get("getCallerIDFiltering", values, rs); err != nil {
		return nil, err
	}

	return rs.CallerIDFilters, nil
}

func (d *DIDsAPI) GetCarriers(carrier string) ([]Carrier, error) {
	values := url.Values{}

	if carrier != "" {
		values.Add("carrier", carrier)
	}

	rs := &GetCarriersResp{}
	if err := d.client.Get("getCarriers", values, rs); err != nil {
		return nil, err
	}

	return rs.Carriers, nil
}

func (d *DIDsAPI) GetDIDCountries(countryId, typ3 string) ([]DIDCountries, error) {
	values := url.Values{}
	values.Add("type", typ3)

	if countryId != "" {
		values.Add("country_id", countryId)
	}

	rs := &GetDIDCountriesResp{}
	if err := d.client.Get("getDIDCountries", values, rs); err != nil {
		return nil, err
	}

	return rs.Countries, nil
}

func (d *DIDsAPI) GetDIDsCan(province, rateCenter string) ([]DID, error) {
	values := url.Values{}
	values.Add("province", province)

	if rateCenter != "" {
		values.Add("ratecenter", rateCenter)
	}

	rs := &GetDIDsCanResp{}
	if err := d.client.Get("getDIDsCAN", values, rs); err != nil {
		return nil, err
	}

	return rs.DIDs, nil
}

func (d *DIDsAPI) GetDIDsInfo(client, DID string) ([]DIDInfo, error) {
	values := url.Values{}

	if client != "" {
		values.Add("client", client)
	}

	if DID != "" {
		values.Add("did", DID)
	}

	rs := &GetDIDsInfoResp{}
	if err := d.client.Get("getDIDsInfo", values, rs); err != nil {
		return nil, err
	}

	return rs.DIDInfos, nil
}

func (d *DIDsAPI) GetDIDsInternationalGeographic(countryId string) ([]InternationalLocations, error) {
	values := url.Values{}
	values.Add("country_id", countryId)

	rs := &GetDIDsInternationalResp{}
	if err := d.client.Get("getDIDsInternationalGeographic", values, rs); err != nil {
		return nil, err
	}

	return rs.Locations, nil
}

func (d *DIDsAPI) GetDIDsInternationalNational(countryId string) ([]InternationalLocations, error) {
	values := url.Values{}
	values.Add("country_id", countryId)

	rs := &GetDIDsInternationalResp{}
	if err := d.client.Get("getDIDsInternationalNational", values, rs); err != nil {
		return nil, err
	}

	return rs.Locations, nil
}

func (d *DIDsAPI) GetDIDsInternationalTollFree(countryId string) ([]InternationalLocations, error) {
	values := url.Values{}
	values.Add("country_id", countryId)

	rs := &GetDIDsInternationalResp{}
	if err := d.client.Get("getDIDsInternationalTollFree", values, rs); err != nil {
		return nil, err
	}

	return rs.Locations, nil
}

func (d *DIDsAPI) GetDIDsUSA(state, rateCenter string) ([]DID, error) {
	values := url.Values{}
	values.Add("state", state)

	if rateCenter != "" {
		values.Add("ratecenter", rateCenter)
	}

	rs := &GetDIDsUSAResp{}
	if err := d.client.Get("getDIDsUSA", values, rs); err != nil {
		return nil, err
	}

	return rs.DIDs, nil
}

func (d *DIDsAPI) GetDISAs(DISA string) ([]DISA, error) {
	values := url.Values{}

	if DISA != "" {
		values.Add("disa", DISA)
	}

	rs := &GetDISAsResp{}
	if err := d.client.Get("getDISAs", values, rs); err != nil {
		return nil, err
	}

	return rs.DISAs, nil
}

func (d *DIDsAPI) GetForwardings(forwarding string) ([]Forwarding, error) {
	values := url.Values{}

	if forwarding != "" {
		values.Add("forwarding", forwarding)
	}

	rs := &GetForwardingsResp{}
	if err := d.client.Get("getForwardings", values, rs); err != nil {
		return nil, err
	}

	return rs.Forwardings, nil
}

func (d *DIDsAPI) GetInternationalTypes(typ3 string) ([]InternationalTypes, error) {
	values := url.Values{}

	if typ3 != "" {
		values.Add("type", typ3)
	}

	rs := &GetInternationalTypesResp{}
	if err := d.client.Get("getInternationalTypes", values, rs); err != nil {
		return nil, err
	}

	return rs.Types, nil
}

func (d *DIDsAPI) GetIVRs(IVR string) ([]IVR, error) {
	values := url.Values{}

	if IVR != "" {
		values.Add("ivr", IVR)
	}

	rs := &GetIVRsResp{}
	if err := d.client.Get("getIVRs", values, rs); err != nil {
		return nil, err
	}

	return rs.IVRs, nil
}

func (d *DIDsAPI) GetJoinWhenEmptyTypes(typ3 string) ([]JoinWhenEmptyType, error) {
	values := url.Values{}

	if typ3 != "" {
		values.Add("type", typ3)
	}

	rs := &GetJoinWhenEmptyTypesResp{}
	if err := d.client.Get("getJoinWhenEmptyTypes", values, rs); err != nil {
		return nil, err
	}

	return rs.JoinWhenEmptyTypes, nil
}

func (d *DIDsAPI) GetPhonebook(phonebook, name string) ([]Phonebook, error) {
	values := url.Values{}

	if phonebook != "" {
		values.Add("phonebook", phonebook)
	}

	if name != "" {
		values.Add("name", name)
	}

	rs := &GetPhonebookResp{}
	if err := d.client.Get("getPhonebook", values, rs); err != nil {
		return nil, err
	}

	return rs.Phonebooks, nil
}

func (d *DIDsAPI) GetPortability(DID string) (bool, []Plan, error) {
	values := url.Values{}
	values.Add("did", DID)

	rs := &GetPortabilityResp{}
	if err := d.client.Get("getPortability", values, rs); err != nil {
		return false, nil, err
	}

	portable := rs.Portable == "yes"

	return portable, rs.Plans, nil
}

func (d *DIDsAPI) GetProvinces() ([]Province, error) {
	values := url.Values{}

	rs := &GetProvincesResp{}
	if err := d.client.Get("getProvinces", values, rs); err != nil {
		return nil, err
	}

	return rs.Provinces, nil
}

func (d *DIDsAPI) GetQueues() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetRateCentersCan(province string) ([]RateCenter, error) {
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

func (d *DIDsAPI) GetRecordings() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetRecordingFile() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetRingGroups() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetRingStrategies() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetSIPURIs() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetSMS() error {
	return errors.New("NOT IMPLEMENTED YET!")
}

func (d *DIDsAPI) GetStates() ([]State, error) {
	rs := &GetStatesResp{}
	if err := d.client.Get("getStates", url.Values{}, rs); err != nil {
		return nil, err
	}

	return rs.States, nil
}

func (d *DIDsAPI) GetStaticMembers(queue, member string) ([]Member, error) {
	values := url.Values{}
	values.Add("queue", queue)

	if member != "" {
		values.Add("member", member)
	}

	rs := &GetStaticMembersResp{}
	if err := d.client.Get("getStaticMembers", values, rs); err != nil {
		return nil, err
	}

	return rs.Members, nil
}

func (d *DIDsAPI) GetTimeConditions(timeCondition string) ([]TimeCondition, error) {
	values := url.Values{}

	if timeCondition != "" {
		values.Add("timecondition", timeCondition)
	}

	rs := &GetTimeConditionsResp{}
	if err := d.client.Get("getTimeConditions", values, rs); err != nil {
		return nil, err
	}

	return rs.TimeConditions, nil
}

func (d *DIDsAPI) GetVoicemailSetups(voicemailSetup string) ([]VoicemailSetup, error) {
	values := url.Values{}

	if voicemailSetup != "" {
		values.Add("voicemailsetup", voicemailSetup)
	}

	rs := &GetVoicemailSetups{}
	if err := d.client.Get("getVoicemailSetups", values, rs); err != nil {
		return nil, err
	}

	return rs.VoicemailSetups, nil
}

func (d *DIDsAPI) GetVoicemailAttachmentFormats(emailAttachmentFormat string) ([]VoicemailAttachmentFormat, error) {
	values := url.Values{}

	if emailAttachmentFormat != "" {
		values.Add("email_attachment_format", emailAttachmentFormat)
	}

	rs := &GetVoicemailAttachmentFormats{}
	if err := d.client.Get("getVoicemailAttachmentFormats", values, rs); err != nil {
		return nil, err
	}

	return rs.VoicemailAttachmentFormats, nil
}

func (d *DIDsAPI) OrderDID(didOrder *DIDOrder) error {
	rs := &BaseResp{}
	rq := didOrder

	if err := d.client.Post("orderDID", rq, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) OrderDIDInternationalGeographic(didOrder *DIDOrderInternationalGeographic) error {
	rs := &BaseResp{}
	rq := didOrder

	if err := d.client.Post("orderDIDInternationalGeographic", rq, rs); err != nil {
		return err
	}

	return nil
}

func (d *DIDsAPI) SearchDIDsCan(province string, typ3 DIDSearchType, query string) ([]DID, error) {
	values := url.Values{}
	values.Add("type", string(typ3))
	values.Add("query", query)

	if province != "" {
		values.Add("province", province)
	}

	rs := &SearchDIDsCanResp{}
	if err := d.client.Get("searchDIDsCAN", values, rs); err != nil {
		return nil, err
	}

	return rs.DIDs, nil
}