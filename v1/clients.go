package v1

import (
	"net/url"
	"fmt"
	"time"
	"encoding/json"
)

type ClientsAPI struct {
	client *VOIPClient
}

type AddChargeResp struct {
	BaseResp
}

type AddChargeReq struct {
	Client      string `json:"client"`
	Charge      string `json:"charge"`
	Description string `json:"description"`
	Test        string `json:"test"`
}

type AddPaymentReq struct {
	Client      string `json:"client"`
	Payment     string `json:"payment"`
	Description string `json:"description"`
	Test        string `json:"test"`
}

type GetBalanceMangementResp struct {
	BaseResp
	BalanceManagement []BalanceManagement `json:"balance_mangement"`
}

type BalanceManagement NumberValueDescription

type GetChargesResp struct {
	BaseResp
	Charges []Charge `json:"charges"`
}

type Charge struct {
	Id          string `json:"id"`
	Date        time.Time `json:"date"`
	Amount      float64 `json:"amount"`
	Description string `json:"description"`
}

func (c *Charge) UnmarshalJSON(data []byte) error {

	type Alias Charge
	aux := &struct {
		Date string `json:"date"`
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

	return nil
}

type GetClientPackagesResp struct {
	BaseResp
	ClientPackages []ClientPackage `json:"packages"`
}

type ClientPackage StringValueDescription

type GetClientsResp struct {
	BaseResp
	Clients []Client `json:"clients"`
}

type Client struct {
	Client            string `json:"client"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Company           string `json:"company,omitempty"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Address           string `json:"address,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Country           string `json:"country,omitempty"`
	Zip               string `json:"zip,omitempty"`
	PhoneNumber       string `json:"phone_number"`
	BalanceManagement string `json:"balance_management,omitempty"`
}

type GetClientThresholdResp struct {
	BaseResp
	ThresholdInformation ClientThreshold `json:"threshold_information"`
}

type ClientThreshold struct {
	Threshold string `json:"threshold"`
	Email     string `json:"email"`
}

type GetDepositsResp struct {
	BaseResp
	Deposits []Deposit `json:"deposits"`
}

type Deposit Charge

type GetPackagesResp struct {
	BaseResp
	Packages []Package `json:"packages"`
}

type Package struct {
	Package            string `json:"package"` //This is the package id.
	Name               string `json:"name"`
	MarkupFixed        string `json:"markup_fixed"`
	MarkupPercentage   string `json:"markup_percentage"`
	Pulse              string `json:"pulse"`
	InternationalRoute string `json:"international_route"`
	CanadaRoute        string `json:"canada_route"`
	MonthlyFee         string `json:"monthly_fee"`
	SetupFee           string `json:"setup_fee"`
	FreeMinutes        string `json:"free_minutes"`
}

type GetResellerBalanceResp struct {
	BaseResp
	Balance Balance `json:"balance"`
}

type SetClientReq Client

//TODO:Stan Req objects aren't needed out side of the package. Change.
type SetClientThresholdReq struct {
	Client    string `json:"client"`
	Threshold string `json:"threshold"`
	Email     string `json:"email"`
}

type SignupClientReq struct {
	Client
	ConfirmEmail    string `json:"confirm_email"`
	ConfirmPassword string `json:"confirm_password"`
	Activate        bool `json:"activate"`
}

func (c *ClientsAPI) AddCharge(client, description string, charge float64, test bool) error {
	rs := &BaseResp{}
	rq := &AddChargeReq{
		Client: client,
		Charge: fmt.Sprintf("%f", charge),
		Description: description,
		Test: fmt.Sprintf("%t", test),
	}

	if err := c.client.Post("addCharge", rq, rs); err != nil {
		return err
	}

	return nil
}

func (c *ClientsAPI) AddPayment(client, description string, payment float64, test bool) error {
	rs := &BaseResp{}
	rq := &AddPaymentReq{
		Client: client,
		Payment: fmt.Sprintf("%f", payment),
		Description: description,
		Test: fmt.Sprintf("%t", test),
	}

	if err := c.client.Post("addPayment", rq, rs); err != nil {
		return err
	}

	return nil
}

func (c *ClientsAPI) GetBalanceManagement(balanceManagement string) ([]BalanceManagement, error) {
	values := url.Values{}
	if balanceManagement != "" {
		values.Add("balance_management", balanceManagement)
	}

	rs := &GetBalanceMangementResp{}
	if err := c.client.Get("getBalanceManagement", values, rs); err != nil {
		return nil, err
	}

	return rs.BalanceManagement, nil
}

func (c *ClientsAPI) GetCharges(client string) ([]Charge, error) {
	values := url.Values{}
	values.Add("client", client)

	rs := &GetChargesResp{}
	if err := c.client.Get("getCharges", values, rs); err != nil {
		return nil, err
	}

	return rs.Charges, nil
}

func (c *ClientsAPI) GetClientPackages(client string) ([]ClientPackage, error) {
	values := url.Values{}
	values.Add("client", client)

	rs := &GetClientPackagesResp{}
	if err := c.client.Get("getClientPackages", values, rs); err != nil {
		return nil, err
	}

	return rs.ClientPackages, nil
}

func (c *ClientsAPI) GetClients(client string) ([]Client, error) {
	values := url.Values{}
	if client != "" {
		values.Add("client", client)
	}

	rs := &GetClientsResp{}
	if err := c.client.Get("getClients", values, rs); err != nil {
		return nil, err
	}

	return rs.Clients, nil
}

func (c *ClientsAPI) GetClientThreshold(client string) (*ClientThreshold, error) {
	values := url.Values{}
	values.Add("client", client)

	rs := &GetClientThresholdResp{}
	if err := c.client.Get("getClientThreshold", values, rs); err != nil {
		return nil, err
	}

	return &rs.ThresholdInformation, nil
}

func (c *ClientsAPI) GetDeposits(client string) ([]Deposit, error) {
	values := url.Values{}
	values.Add("client", client)

	rs := &GetDepositsResp{}
	if err := c.client.Get("getDeposits", values, rs); err != nil {
		return nil, err
	}

	return rs.Deposits, nil
}

func (c *ClientsAPI) GetPackages(packag3 string) ([]Package, error) {
	values := url.Values{}
	if packag3 != "" {
		values.Add("package", packag3)
	}

	rs := &GetPackagesResp{}
	if err := c.client.Get("getPackages", values, rs); err != nil {
		return nil, err
	}

	return rs.Packages, nil
}

func (c *ClientsAPI) GetResellerBalance(client string) (*Balance, error) {
	values := url.Values{}
	values.Add("client", client)

	rs := &GetResellerBalanceResp{}
	if err := c.client.Get("getResellerBalance", values, rs); err != nil {
		return nil, err
	}

	return &rs.Balance, nil
}

func (c *ClientsAPI) SetClient(client *Client) error {
	rs := &BaseResp{}
	rq := *client

	if err := c.client.Post("setClient", rq, rs); err != nil {
		return err
	}

	return nil
}

func (c *ClientsAPI) SetClientThreshold(client, threshold, email string) error {
	rs := &BaseResp{}
	rq := &SetClientThresholdReq{
		client,
		threshold,
		email,
	}

	if err := c.client.Post("setClientThreshold", rq, rs); err != nil {
		return err
	}

	return nil
}

func (c *ClientsAPI) SignupClient(client *Client, confirmEmail, confirmPassword string, activate bool) error {
	rs := &BaseResp{}
	rq := &SignupClientReq{
		*client,
		confirmEmail,
		confirmPassword,
		activate,
	}

	if err := c.client.Post("signupClient", rq, rs); err != nil {
		return err
	}

	return nil
}