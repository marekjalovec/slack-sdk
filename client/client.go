package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const BaseApiUrl = "https://slack.com/api"

type Client struct {
	SigningSecret string
	BotToken      string
}

func (at *Client) validateEvent(body string, timestamp string, signature string) error {
	var mac = hmac.New(sha256.New, []byte(at.SigningSecret))
	mac.Write([]byte(fmt.Sprintf("v0:%s:%s", timestamp, body)))
	var expectedMAC = mac.Sum(nil)
	var expectedHEX = "v0=" + hex.EncodeToString(expectedMAC)

	var signatureValid = hmac.Equal([]byte(signature), []byte(expectedHEX))
	if !signatureValid {
		return fmt.Errorf("event signature not valid")
	}

	return nil
}

func (at *Client) createAuthorizedRequest(method string, endpoint string, body io.Reader) *http.Request {
	var req, _ = http.NewRequest(method, at.getUrl(endpoint), body)

	// set headers and query params
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", at.BotToken))

	return req
}

func (at *Client) Get(endpoint string, payload map[string]string) ([]byte, error) {
	var req = at.createAuthorizedRequest(http.MethodGet, endpoint, nil)
	req.URL.RawQuery = at.getQuery(payload)

	log.Println(fmt.Sprintf("[SLACK]:[GET] Calling %s with params %s", at.getUrl(endpoint), req.URL.RawQuery))

	body, err := at.do(req)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (at *Client) Post(endpoint string, payload interface{}) ([]byte, error) {
	var req = at.createAuthorizedRequest(http.MethodPost, endpoint, at.getBody(payload))

	var p, _ = json.Marshal(payload)
	log.Println(fmt.Sprintf("[SLACK]:[POST] Calling %s with params %s", at.getUrl(endpoint), p))

	body, err := at.do(req)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (at *Client) do(r *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("[SLACK]:[%s] Error: %s", r.Method, err))
		return []byte{}, err
	}
	log.Println(fmt.Sprintf("[SLACK]:[%s] %s", r.Method, resp.Status))

	return body, nil
}

func (at *Client) getUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", BaseApiUrl, endpoint)
}

func (at *Client) getBody(payload interface{}) io.Reader {
	var p, _ = json.Marshal(payload)

	return strings.NewReader(string(p))
}

func (at *Client) getQuery(params map[string]string) string {
	data := url.Values{}
	for key, value := range params {
		data.Add(key, value)
	}

	return data.Encode()
}
