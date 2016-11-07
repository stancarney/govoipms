package v1

import "net/url"

type ClientAPI struct {
	client *VOIPClient
}

type GetPackagesResp struct {
	BaseResp
	Packages []Package `json:"packages"`
}

type Package struct {
	Package            string `json:"package"`
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

func (c *ClientAPI) GetPackages(packag3 string) ([]Package, error) {
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