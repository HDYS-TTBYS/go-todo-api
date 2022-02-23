package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(NewHealthController().Get))
	defer testServer.Close()

	res, err := http.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	respBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, string(respBody), `{"status":200,"message":"OK"}`)

}
