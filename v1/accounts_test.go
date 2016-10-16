package v1

import (
	"github.com/stretchr/testify/require"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"encoding/json"
)

func TestAccountsAPI_GetAllowedCodecs(t *testing.T) {
	
	//setup
	rq := GetAllowedCodecsResp{
		BaseResp{"success"},
		[]Codec{{"ulaw", "test"},{"911", "test1"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func TestAccountsAPI_GetAllowedCodecs_Specified(t *testing.T) {

	//setup
	rq := GetAllowedCodecsResp{
		BaseResp{"success"},
		[]Codec{{"911", "test1"}},
	}
	result, _ := json.Marshal(rq)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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