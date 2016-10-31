package v1

import (
	"testing"
	"encoding/json"
	"net/http/httptest"
	"github.com/stretchr/testify/require"
	"fmt"
	"net/http"
)

func TestGeneralAPI_GetBalance_False(t *testing.T) {

	//setup
	rq := GetBalanceResp {
		BaseResp{"success"},
		Balance{CurrentBalance: "100"},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getBalance"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	balance, err := api.GetBalance(false)

	//verify
	require.NoError(t, err)
	require.Equal(t, rq.Balance, *balance)
}

func TestGeneralAPI_GetBalance_Error(t *testing.T) {

	//setup
	rq := GetBalanceResp {
		BaseResp{"error"},
		Balance{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getBalance"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	balance, err := api.GetBalance(false)

	//verify
	require.Nil(t, balance)
	require.EqualError(t, err, "error")
}

func TestGeneralAPI_GetBalance_True(t *testing.T) {

	//setup
	rq := GetBalanceResp {
		BaseResp{"success"},
		Balance{"100", 0.0, 0.1, "60", 0.2, 1, "55"},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getBalance"}, r.URL.Query()["method"])
		require.Equal(t, []string{"true"}, r.URL.Query()["advanced"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	balance, err := api.GetBalance(true)

	//verify
	require.NoError(t, err)
	require.Equal(t, rq.Balance, *balance)
}

func TestGeneralAPI_GetCountries(t *testing.T) {

	//setup
	rq := GetCountriesResp {
		BaseResp{"success"},
		[]Country{{"CA", "Canada"}, {"US", "United States"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCountries"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	countries, err := api.GetCountries("")

	//verify
	require.NoError(t, err)
	require.Len(t, countries, 2)
	require.Equal(t, rq.Countries, countries)
}

func TestGeneralAPI_GetCountries_Error(t *testing.T) {

	//setup
	rq := GetCountriesResp {
		BaseResp{"error"},
		[]Country{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCountries"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	countries, err := api.GetCountries("")

	//verify
	require.Nil(t, countries)
	require.EqualError(t, err, "error")
}

func TestGeneralAPI_GetCountries_Specified(t *testing.T) {

	//setup
	rq := GetCountriesResp {
		BaseResp{"success"},
		[]Country{{"CA", "Canada"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCountries"}, r.URL.Query()["method"])
		require.Equal(t, []string{"CA"}, r.URL.Query()["country"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	countries, err := api.GetCountries("CA")

	//verify
	require.NoError(t, err)
	require.Len(t, countries, 1)
	require.Equal(t, rq.Countries, countries)
}

func TestGeneralAPI_GetIP(t *testing.T) {

	//setup
	rq := GetIPResp  {
		BaseResp{"success"},
		"127.0.0.1",
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getIP"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	ip, err := api.GetIP()

	//verify
	require.NoError(t, err)
	require.Equal(t, rq.IP, ip)
}

func TestGeneralAPI_GetIP_Error(t *testing.T) {

	//setup
	rq := GetIPResp  {
		BaseResp{"error"},
		"",
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getIP"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	ip, err := api.GetIP()

	//verify
	require.Equal(t, ip, "")
	require.EqualError(t, err, "error")
}

func TestGeneralAPI_GetLanguages(t *testing.T) {

	//setup
	rq := GetLanguagesResp {
		BaseResp{"success"},
		[]Language{{"en", "English"}, {"fr", "French"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLanguages"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	languages, err := api.GetLanguages("")

	//verify
	require.NoError(t, err)
	require.Len(t, languages, 2)
	require.Equal(t, rq.Languages, languages)
}

func TestGeneralAPI_GetLanguages_Error(t *testing.T) {

	//setup
	rq := GetLanguagesResp {
		BaseResp{"error"},
		[]Language{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLanguages"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	languages, err := api.GetLanguages("")

	//verify
	require.Nil(t, languages)
	require.EqualError(t, err, "error")
}

func TestGeneralAPI_GetLanguages_Specified(t *testing.T) {

	//setup
	rq := GetLanguagesResp {
		BaseResp{"success"},
		[]Language{{"en", "English"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLanguages"}, r.URL.Query()["method"])
		require.Equal(t, []string{"en"}, r.URL.Query()["language"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	languages, err := api.GetLanguages("en")

	//verify
	require.NoError(t, err)
	require.Len(t, languages, 1)
	require.Equal(t, rq.Languages, languages)
}

func TestGeneralAPI_GetServerInfo(t *testing.T) {

	//setup
	rq := GetServerInfoResp {
		BaseResp{"success"},
		[]Server{{"ServerName1", "ServerShortname1", "ServerHostname1", "ServerIP1", "ServerCountry1", "ServerPOP1"},
			{"ServerName2", "ServerShortname2", "ServerHostname2", "ServerIP2", "ServerCountry2", "ServerPOP2"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getServersInfo"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	servers, err := api.GetServerInfo("")

	//verify
	require.NoError(t, err)
	require.Len(t, servers, 2)
	require.Equal(t, rq.Servers, servers)
}

func TestGeneralAPI_GetServerInfo_Error(t *testing.T) {

	//setup
	rq := GetServerInfoResp {
		BaseResp{"error"},
		[]Server{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getServersInfo"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	servers, err := api.GetServerInfo("")

	//verify
	require.Nil(t, servers)
	require.EqualError(t, err, "error")
}

func TestGeneralAPI_GetServerInfo_Specified(t *testing.T) {

	//setup
	rq := GetServerInfoResp {
		BaseResp{"success"},
		[]Server{{"ServerName1", "ServerShortname1", "ServerHostname1", "ServerIP1", "ServerCountry1", "ServerPOP1"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getServersInfo"}, r.URL.Query()["method"])
		require.Equal(t, []string{"ServerName1"}, r.URL.Query()["server_pop"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewGeneralAPI()

	//execute
	servers, err := api.GetServerInfo("ServerName1")

	//verify
	require.NoError(t, err)
	require.Len(t, servers, 1)
	require.Equal(t, rq.Servers, servers)
}


