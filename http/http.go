package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpRequest struct {
	Client  *http.Client
	Request *http.Request
}

type Response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
}

func NewHttpRequest(method string, url string, token string, data *bytes.Buffer) (*HttpRequest, error) {

	var req *http.Request

	var err error

	if data != nil {
		req, err = http.NewRequest(method, url, data)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(token) != "" {
		bearer := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", bearer)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	return &HttpRequest{
		Client:  client,
		Request: req,
	}, nil

}

func (h *HttpRequest) Post() (*Response, error) {

	res, err := h.Client.Do(h.Request)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return nil, err
	}

	var rs Response
	json.Unmarshal(data, &rs)

	return &rs, nil

}

func (h *HttpRequest) Execute() ([]byte, error) {

	res, err := h.Client.Do(h.Request)

	if res.StatusCode == 401 {
		return nil, fmt.Errorf("error Unauthorized. Please login to BPAAS", res.Status)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error %s", res.Status)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return nil, err
	}

	// var rs Response
	// json.Unmarshal(data, &rs)

	return data, nil

}
