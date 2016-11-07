package v1

import (
	"github.com/stretchr/testify/require"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
)

func TestCDRAPI_GetCallAccounts(t *testing.T) {

	//setup
	rq := GetCallAccountsResp{
		BaseResp{"success"},
		[]CallAccount{{"all", "All Accounts"}, {"100000_VoIP", "100000_VoIP"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallAccounts"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	accounts, err := api.GetCallAccounts("")

	//verify
	require.NoError(t, err)
	require.Len(t, accounts, 2)
	require.Equal(t, rq.CallAccounts, accounts)
}

func TestCDRAPI_GetCallAccounts_Error(t *testing.T) {

	//setup
	rq := GetCallAccountsResp{
		BaseResp{"error"},
		[]CallAccount{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallAccounts"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	accounts, err := api.GetCallAccounts("")

	//verify
	require.EqualError(t, err, "error")
	require.Len(t, accounts, 0)
}

func TestCDRAPI_GetCallAccounts_Specified(t *testing.T) {

	//setup
	rq := GetCallAccountsResp{
		BaseResp{"success"},
		[]CallAccount{{"100000_VoIP", "100000_VoIP"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallAccounts"}, r.URL.Query()["method"])
		require.Equal(t, []string{"100000_VoIP"}, r.URL.Query()["client"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	accounts, err := api.GetCallAccounts("100000_VoIP")

	//verify
	require.NoError(t, err)
	require.Len(t, accounts, 1)
	require.Equal(t, rq.CallAccounts, accounts)
}

func TestCDRAPI_GetCallBilling(t *testing.T) {

	//setup
	rq := GetCallBillingResp{
		BaseResp{"success"},
		[]CallBilling{{"all", "All Calls"}, {"free", "Free Calls"}, {"billed", "Billed Calls"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallBilling"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	callBillings, err := api.GetCallBilling()

	//verify
	require.NoError(t, err)
	require.Len(t, callBillings, 3)
	require.Equal(t, rq.CallBilling, callBillings)
}

func TestCDRAPI_GetCallBilling_Error(t *testing.T) {

	//setup
	rq := GetCallBillingResp{
		BaseResp{"error"},
		[]CallBilling{{"all", "All Calls"}, {"free", "Free Calls"}, {"billed", "Billed Calls"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallBilling"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	callBillings, err := api.GetCallBilling()

	//verify
	require.EqualError(t, err, "error")
	require.Len(t, callBillings, 0)
}

func TestCDRAPI_GetCallTypes(t *testing.T) {

	//setup
	rq := GetCallTypeResp{
		BaseResp{"success"},
		[]CallType{{"all", "All Calls"}, {"outgoing", "Outgoing Calls"}, {"incoming", "Incoming Calls"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	callTypes, err := api.GetCallTypes("")

	//verify
	require.NoError(t, err)
	require.Len(t, callTypes, 3)
	require.Equal(t, rq.CallTypes, callTypes)
}

func TestCDRAPI_GetCallTypes_Error(t *testing.T) {

	//setup
	rq := GetCallTypeResp{
		BaseResp{"error"},
		[]CallType{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	callTypes, err := api.GetCallTypes("")

	//verify
	require.EqualError(t, err, "error")
	require.Len(t, callTypes, 0)
}

func TestCDRAPI_GetCallTypes_Specified(t *testing.T) {

	//setup
	rq := GetCallTypeResp{
		BaseResp{"success"},
		[]CallType{{"all", "All Calls"}, {"outgoing", "Outgoing Calls"}, {"incoming", "Incoming Calls"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCallTypes"}, r.URL.Query()["method"])
		require.Equal(t, []string{"1234"}, r.URL.Query()["client"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	//execute
	callTypes, err := api.GetCallTypes("1234")

	//verify
	require.NoError(t, err)
	require.Len(t, callTypes, 3)
	require.Equal(t, rq.CallTypes, callTypes)
}

func TestCDRAPI_GetCDR(t *testing.T) {

	//setup
	//Magic string was easier than trying to Marshal a CDR record. Something that the API will never need.
	result := `{"status":"success","cdr":[{
	"date":"2016-11-07 10:17:34",
	"callerid":"\"15554443333\" <15554443333>",
	"destination":"15554441111",
	"description":"Inbound DID",
	"account":"123456",
	"disposition":"ANSWERED",
	"duration":"00:00:05",
	"seconds":"5",
	"rate":"0.00900000",
	"total":"0.00090000",
	"uniqueid":"982384595"
	}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getCDR"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2016-10-26"}, r.URL.Query()["date_from"])
		require.Equal(t, []string{"2016-11-07"}, r.URL.Query()["date_to"])
		require.Equal(t, []string{"1"}, r.URL.Query()["answered"])
		require.Equal(t, []string{"1"}, r.URL.Query()["noanswer"])
		require.Equal(t, []string{"1"}, r.URL.Query()["busy"])
		require.Equal(t, []string{"1"}, r.URL.Query()["failed"])
		require.Equal(t, []string{"-7"}, r.URL.Query()["timezone"])
		require.Equal(t, []string{"all"}, r.URL.Query()["calltype"])
		require.Equal(t, []string{"cb"}, r.URL.Query()["callbilling"])
		require.Equal(t, []string{"a"}, r.URL.Query()["account"])
		fmt.Fprintln(w, result)
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	cs := CallStatus{true, true, true, true}
	mst, _ := time.LoadLocation("America/Edmonton")
	dateFrom := time.Now().Add(time.Hour * -300)
	dateTo := time.Now()

	//execute
	cdrs, err := api.GetCDR(dateFrom, dateTo, cs, mst, "all", "cb", "a")

	//verify
	require.NoError(t, err)
	require.Len(t, cdrs, 1)
	require.Equal(t, "\"15554443333\" <15554443333>", cdrs[0].CallerId)
	require.Equal(t, "15554441111", cdrs[0].Destination)
	require.Equal(t, "Inbound DID", cdrs[0].Description)
	require.Equal(t, "123456", cdrs[0].Account)
	require.Equal(t, "ANSWERED", cdrs[0].Disposition)
	require.Equal(t, time.Second * 5, cdrs[0].Duration)
	require.Equal(t, 0.009, cdrs[0].Rate)
	require.Equal(t, 0.0009, cdrs[0].Total)
	require.Equal(t, "982384595", cdrs[0].UniqueId)
}

func TestCDRAPI_GetCDR_Error(t *testing.T) {

	//setup
	//Magic string was easier than trying to Marshal a CDR record. Something that the API will never need.
	result := `{"status":"invalid_something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, result)
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	cs := CallStatus{true, false, false, false}
	mst, _ := time.LoadLocation("America/Edmonton")
	dateFrom := time.Now().Add(time.Hour * -300)
	dateTo := time.Now()

	//execute
	cdrs, err := api.GetCDR(dateFrom, dateTo, cs, mst, "all", "", "")

	//verify
	require.Error(t, err)
	require.Len(t, cdrs, 0)
}

func TestCDRAPI_GetCDR_NoFromDate(t *testing.T) {

	//setup
	//Magic string was easier than trying to Marshal a CDR record. Something that the API will never need.
	result := `{"status":"error"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, result)
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	cs := CallStatus{true, false, false, false}
	mst, _ := time.LoadLocation("America/Edmonton")
	dateFrom := time.Time{}
	dateTo := time.Now()

	//execute
	cdrs, err := api.GetCDR(dateFrom, dateTo, cs, mst, "all", "", "")

	//verify
	require.Error(t, err)
	require.EqualError(t, err, "dateFrom is required!")
	require.Len(t, cdrs, 0)
}

func TestCDRAPI_GetCDR_NoDateTo(t *testing.T) {

	//setup
	//Magic string was easier than trying to Marshal a CDR record. Something that the API will never need.
	result := `{"status":"error"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, result)
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	cs := CallStatus{true, false, false, false}
	mst, _ := time.LoadLocation("America/Edmonton")
	dateFrom := time.Now().Add(time.Hour * -300)
	dateTo := time.Time{}

	//execute
	cdrs, err := api.GetCDR(dateFrom, dateTo, cs, mst, "all", "", "")

	//verify
	require.Error(t, err)
	require.EqualError(t, err, "dateTo is required!")
	require.Len(t, cdrs, 0)
}

func TestCDRAPI_GetCDR_NoTimezone(t *testing.T) {

	//setup
	//Magic string was easier than trying to Marshal a CDR record. Something that the API will never need.
	result := `{"status":"error"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, result)
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewCDRAPI()

	cs := CallStatus{true, false, false, false}
	mst := &time.Location{}
	dateFrom := time.Now().Add(time.Hour * -300)
	dateTo := time.Now()

	//execute
	cdrs, err := api.GetCDR(dateFrom, dateTo, cs, mst, "all", "", "")

	//verify
	require.Error(t, err)
	require.EqualError(t, err, "timezone is required!")
	require.Len(t, cdrs, 0)
}
