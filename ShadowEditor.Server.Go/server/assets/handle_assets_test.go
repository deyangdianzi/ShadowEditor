package assets

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tengge1/shadoweditor/server"
)

func TestAssetsList(t *testing.T) {
	server.Create("../config.toml")

	assets := Assets{}

	ts := httptest.NewServer(http.HandlerFunc(assets.List))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(bytes))
}
