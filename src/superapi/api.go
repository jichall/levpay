package superapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	token string
)

const (
	searchID    = `https://superheroapi.com/api/%s/%s`
	searchName  = `https://superheroapi.com/api/%s/search/%s`
	powerstats  = `https://superheroapi.com/api/%s/%s/powerstats`
	connections = `https://superheroapi.com/api/%s/%s/connections`
)

// Fetch returns a response from SuperAPI given a name or an ID
func Fetch(param interface{}) *Response {
	url := ""
	// first, do a type check to see what url we are going to use
	switch param.(type) {
	case string:
		url = fmt.Sprint(searchName, token, param.(string))
	case int:
		url = fmt.Sprint(searchID, token, param.(int))
	}

	res, err := http.Get(url)

	if err != nil {
		logger.Error("failed to request URL %s, reason %v", url, err)
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		logger.Error("failed to read response body, reason %v", err)
		return nil
	}

	resp := &Response{}
	// what if the response is an error?
	err = json.Unmarshal(b, resp)
	if err != nil {
		logger.Error("failed to unmarshal response, reason %v", err)
		return nil
	}

	return resp
}

func GetConnections() {}
func GetPower()       {}
