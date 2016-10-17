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

func TestAccountsAPI_GetDeviceTypes(t *testing.T) {

	//setup
	rq := GetDeviceTypesResp{
		BaseResp{"success"},
		[]DeviceType{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDeviceTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	deviceTypes, err := api.GetDeviceTypes(0)

	//verify
	require.NoError(t, err)
	require.Len(t, deviceTypes, 2)
	require.Equal(t, rq.DeviceTypes, deviceTypes)
}

func TestAccountsAPI_GetDeviceTypes_Error(t *testing.T) {

	//setup
	rq := GetDeviceTypesResp{
		BaseResp{"error"},
		[]DeviceType{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDeviceTypes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	authTypes, err := api.GetDeviceTypes(0)

	//verify
	require.Nil(t, authTypes)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetDeviceTypes_Specified(t *testing.T) {

	//setup
	rq := GetDeviceTypesResp{
		BaseResp{"success"},
		[]DeviceType{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDeviceTypes"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["device_type"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	deviceTypes, err := api.GetDeviceTypes(2)

	//verify
	require.NoError(t, err)
	require.Len(t, deviceTypes, 1)
	require.Equal(t, rq.DeviceTypes[0], deviceTypes[0])
}

func TestAccountsAPI_GetDTMFModes(t *testing.T) {

	//setup
	rq := GetDTMFModesResp{
		BaseResp{"success"},
		[]DTMFMode{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDTMFModes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	dtmfModes, err := api.GetDTMFModes("")

	//verify
	require.NoError(t, err)
	require.Len(t, dtmfModes, 2)
	require.Equal(t, rq.DTMFModes, dtmfModes)
}

func TestAccountsAPI_GetDTMFModes_Error(t *testing.T) {

	//setup
	rq := GetDTMFModesResp{
		BaseResp{"error"},
		[]DTMFMode{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDTMFModes"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	dtmfModes, err := api.GetDTMFModes("")

	//verify
	require.Nil(t, dtmfModes)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetDTMFModes_Specified(t *testing.T) {

	//setup
	rq := GetDTMFModesResp{
		BaseResp{"success"},
		[]DTMFMode{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getDTMFModes"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["dtmf_mode"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	dtmfModes, err := api.GetDTMFModes("2")

	//verify
	require.NoError(t, err)
	require.Len(t, dtmfModes, 1)
	require.Equal(t, rq.DTMFModes[0], dtmfModes[0])
}

func TestAccountsAPI_GetLockInternational(t *testing.T) {

	//setup
	rq := GetLockInternationalResp{
		BaseResp{"success"},
		[]LockInternational{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLockInternational"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	lockInternational, err := api.GetLockInternational("")

	//verify
	require.NoError(t, err)
	require.Len(t, lockInternational, 2)
	require.Equal(t, rq.LockInternational, lockInternational)
}

func TestAccountsAPI_GetLockInternational_Error(t *testing.T) {

	//setup
	rq := GetLockInternationalResp{
		BaseResp{"error"},
		[]LockInternational{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLockInternational"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	lockInternational, err := api.GetLockInternational("")

	//verify
	require.Nil(t, lockInternational)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetLockInternational_Specified(t *testing.T) {

	//setup
	rq := GetLockInternationalResp{
		BaseResp{"success"},
		[]LockInternational{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getLockInternational"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["lock_international"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	lockInternational, err := api.GetLockInternational("2")

	//verify
	require.NoError(t, err)
	require.Len(t, lockInternational, 1)
	require.Equal(t, rq.LockInternational[0], lockInternational[0])
}

func TestAccountsAPI_GetMusicOnHold(t *testing.T) {

	//setup
	rq := GetMusicOnHoldResp{
		BaseResp{"success"},
		[]MusicOnHold{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getMusicOnHold"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	musicOnHold, err := api.GetMusicOnHold("")

	//verify
	require.NoError(t, err)
	require.Len(t, musicOnHold, 2)
	require.Equal(t, rq.MusicOnHold, musicOnHold)
}

func TestAccountsAPI_GetMusicOnHold_Error(t *testing.T) {

	//setup
	rq := GetMusicOnHoldResp{
		BaseResp{"error"},
		[]MusicOnHold{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getMusicOnHold"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	musicOnHold, err := api.GetMusicOnHold("")

	//verify
	require.Nil(t, musicOnHold)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetMusicOnHold_Specified(t *testing.T) {

	//setup
	rq := GetMusicOnHoldResp{
		BaseResp{"success"},
		[]MusicOnHold{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getMusicOnHold"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["music_on_hold"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	musicOnHold, err := api.GetMusicOnHold("2")

	//verify
	require.NoError(t, err)
	require.Len(t, musicOnHold, 1)
	require.Equal(t, rq.MusicOnHold[0], musicOnHold[0])
}

func TestAccountsAPI_GetNAT(t *testing.T) {

	//setup
	rq := GetNATResp{
		BaseResp{"success"},
		[]NAT{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getNAT"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	NAT, err := api.GetNAT("")

	//verify
	require.NoError(t, err)
	require.Len(t, NAT, 2)
	require.Equal(t, rq.NAT, NAT)
}

func TestAccountsAPI_GetNAT_Error(t *testing.T) {

	//setup
	rq := GetNATResp{
		BaseResp{"error"},
		[]NAT{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getNAT"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	NAT, err := api.GetNAT("")

	//verify
	require.Nil(t, NAT)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetNAT_Specified(t *testing.T) {

	//setup
	rq := GetNATResp{
		BaseResp{"success"},
		[]NAT{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getNAT"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["nat"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	NAT, err := api.GetNAT("2")

	//verify
	require.NoError(t, err)
	require.Len(t, NAT, 1)
	require.Equal(t, rq.NAT[0], NAT[0])
}

func TestAccountsAPI_GetProtocols(t *testing.T) {

	//setup
	rq := GetProtocolResp{
		BaseResp{"success"},
		[]Protocol{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getProtocols"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	protocols, err := api.GetProtocols(0)

	//verify
	require.NoError(t, err)
	require.Len(t, protocols, 2)
	require.Equal(t, rq.Protocols, protocols)
}

func TestAccountsAPI_GetProtocols_Error(t *testing.T) {

	//setup
	rq := GetProtocolResp{
		BaseResp{"error"},
		[]Protocol{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getProtocols"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	protocols, err := api.GetProtocols(0)

	//verify
	require.Nil(t, protocols)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetProtocols_Specified(t *testing.T) {

	//setup
	rq := GetProtocolResp{
		BaseResp{"success"},
		[]Protocol{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getProtocols"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["protocol"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	protocols, err := api.GetProtocols(2)

	//verify
	require.NoError(t, err)
	require.Len(t, protocols, 1)
	require.Equal(t, rq.Protocols[0], protocols[0])
}

func TestAccountsAPI_GetRegistrationStatus_NotSpecified(t *testing.T) {

	//setup
	rq := GetRegistrationStatusResp{
		BaseResp{"error"},
		"no",
		[]RegistrationStatus{{}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getRegistrationStatus"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	registered, registrations, err := api.GetRegistrationStatus("")

	//verify
	require.Nil(t, registrations)
	require.False(t, registered)
	require.EqualError(t, err, "missing_account")
}

func TestAccountsAPI_GetRegistrationStatus_Error(t *testing.T) {

	//setup
	rq := GetRegistrationStatusResp{
		BaseResp{"error"},
		"no",
		[]RegistrationStatus{{}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getRegistrationStatus"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	registered, registrations, err := api.GetRegistrationStatus("2")

	//verify
	require.Nil(t, registrations)
	require.False(t, registered)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetRegistrationStatus_Specified(t *testing.T) {

	//setup
	rq := GetRegistrationStatusResp{
		BaseResp{"success"},
		"yes",
		[]RegistrationStatus{{
			Server{
				"ServerName",
				"ServerShortname",
				"example.com",
				"127.0.0.1",
				"CA",
				"15",
			},
			"192.168.1.1",
			"59870",
			"2010-11-30 16:48:30",
		}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getRegistrationStatus"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["account"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	registered, registrations, err := api.GetRegistrationStatus("2")

	//verify
	require.NoError(t, err)
	require.Len(t, registrations, 1)
	require.True(t, registered)
	require.Equal(t, rq.Registrations, registrations)
}

func TestAccountsAPI_GetReportEstimatedHoldTime(t *testing.T) {

	//setup
	rq := GetReportEstimatedHoldTimeResp{
		BaseResp{"success"},
		[]EstimatedHoldTime{{"1", "one"}, {"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getReportEstimatedHoldTime"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	types, err := api.GetReportEstimatedHoldTime("")

	//verify
	require.NoError(t, err)
	require.Len(t, types, 2)
	require.Equal(t, rq.Types, types)
}

func TestAccountsAPI_GetReportEstimatedHoldTime_Error(t *testing.T) {

	//setup
	rq := GetReportEstimatedHoldTimeResp{
		BaseResp{"error"},
		[]EstimatedHoldTime{},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getReportEstimatedHoldTime"}, r.URL.Query()["method"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	types, err := api.GetReportEstimatedHoldTime("")

	//verify
	require.Nil(t, types)
	require.EqualError(t, err, "error")
}

func TestAccountsAPI_GetReportEstimatedHoldTime_Specified(t *testing.T) {

	//setup
	rq := GetReportEstimatedHoldTimeResp{
		BaseResp{"success"},
		[]EstimatedHoldTime{{"2", "two"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, []string{"getReportEstimatedHoldTime"}, r.URL.Query()["method"])
		require.Equal(t, []string{"2"}, r.URL.Query()["type"])
		fmt.Fprintln(w, string(result))
	}))
	defer ts.Close()

	api := NewVOIPClient(ts.URL, "", "", true).NewAccountsAPI()

	//execute
	types, err := api.GetReportEstimatedHoldTime("2")

	//verify
	require.NoError(t, err)
	require.Len(t, types, 1)
	require.Equal(t, rq.Types[0], types[0])
}