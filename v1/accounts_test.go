package v1

import (
	"github.com/stretchr/testify/require"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

func TestAccountsAPI_CreateSubAccount(t *testing.T) {

	//setup
	rq := CreateSubAccountResp{
		BaseResp{"success"},
		12345,
		"12345_account",
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	a := &Account{
		Username: "Test1",
		Protocol: "1",
		Description: "Description",
		AuthType: "1",
		Password: "Password1",
		IP: "",
		DeviceType: "2",
		CalleridNumber: "5555551234",
		CanadaRouting: "1",
		LockInternational: "1",
		InternationalRoute: "1",
		MusicOnHold: "default",
		AllowedCodecs: "ulaw;g729",
		DTMFMode: "auto",
		NAT: "yes",
		InternalExtension: "",
		InternalVoicemail: "",
		InternalDialtime: "20",
		ResellerClient: "0",
		ResellerPackage: "0",
		ResellerNextbilling: "0000-00-00",
	}

	//execute
	err := api.CreateSubAccount(a)

	//verify
	require.NoError(t, err)
	require.Equal(t, strconv.Itoa(rq.Id), a.Id)
	require.Equal(t, rq.Account, a.Account)
}

func TestAccountsAPI_CreateSubAccount_Error(t *testing.T) {

	//setup
	rq := CreateSubAccountResp{
		BaseResp{"error"},
		0,
		"",
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	a := &Account{
		Username: "Test1",
		Protocol: "1",
		Description: "Description",
		AuthType: "1",
		Password: "Password1",
		IP: "",
		DeviceType: "2",
		CalleridNumber: "5555551234",
		CanadaRouting: "1",
		LockInternational: "1",
		InternationalRoute: "1",
		MusicOnHold: "default",
		AllowedCodecs: "ulaw;g729",
		DTMFMode: "auto",
		NAT: "yes",
		InternalExtension: "",
		InternalVoicemail: "",
		InternalDialtime: "20",
		ResellerClient: "0",
		ResellerPackage: "0",
		ResellerNextbilling: "0000-00-00",
	}

	//execute
	err := api.CreateSubAccount(a)

	//verify
	require.Error(t, err)
	require.Equal(t, "", a.Id)
	require.Equal(t, "", a.Account)
}

func TestAccountsAPI_DelSubAccount(t *testing.T) {

	//setup
	rq := DelSubAccountResp{
		BaseResp{"success"},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	err := api.DelSubAccount("12345")

	//verify
	require.NoError(t, err)
}

func TestAccountsAPI_DelSubAccount_Error(t *testing.T) {

	//setup
	rq := DelSubAccountResp{
		BaseResp{"error"},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	err := api.DelSubAccount("12345")

	//verify
	require.Error(t, err)
}

func TestAccountsAPI_GetAllowedCodecs(t *testing.T) {

	//setup
	rq := GetAllowedCodecsResp{
		BaseResp{"success"},
		[]Codec{{"ulaw", "test"}, {"911", "test1"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAllowedCodecs"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	codecs, err := api.GetAllowedCodecs("")

	//verify
	require.NoError(t, err)
	require.Len(t, codecs, 2)
	require.Equal(t, rq.AllowedCodecs, codecs)
}

func TestAccountsAPI_GetAllowedCodecs_Error(t *testing.T) {

	//setup
	rq := GetAllowedCodecsResp{
		BaseResp{"error"},
		[]Codec{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAllowedCodecs"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	codecs, err := api.GetAllowedCodecs("")

	//verify
	require.Nil(t, codecs)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetAllowedCodecs_Specified(t *testing.T) {

	//setup
	rq := GetAllowedCodecsResp{
		BaseResp{"success"},
		[]Codec{{"911", "test1"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAllowedCodecs"}, r.URL.Query()["method"])
		require.Equal(t, []string{"911"}, r.URL.Query()["codec"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	codecs, err := api.GetAllowedCodecs("911")

	//verify
	require.NoError(t, err)
	require.Len(t, codecs, 1)
	require.Equal(t, rq.AllowedCodecs[0], codecs[0])
}

func TestAccountsAPI_GetAuthTypes(t *testing.T) {

	//setup
	rq := GetAuthTypesResp{
		BaseResp{"success"},
		[]AuthType{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAuthTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	authTypes, err := api.GetAuthTypes(0)

	//verify
	require.NoError(t, err)
	require.Len(t, authTypes, 2)
	require.Equal(t, rq.AuthTypes, authTypes)
}

func TestAccountsAPI_GetAuthTypes_Error(t *testing.T) {

	//setup
	rq := GetAuthTypesResp{
		BaseResp{"error"},
		[]AuthType{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAuthTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	authTypes, err := api.GetAuthTypes(0)

	//verify
	require.Nil(t, authTypes)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetAuthTypes_Specified(t *testing.T) {

	//setup
	rq := GetAuthTypesResp{
		BaseResp{"success"},
		[]AuthType{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getAuthTypes"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["type"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	authTypes, err := api.GetAuthTypes(2)

	//verify
	require.NoError(t, err)
	require.Len(t, authTypes, 1)
	require.Equal(t, rq.AuthTypes[0], authTypes[0])
}