package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default DatoCMS URL
const HostURL string = "https://site-api.datocms.com/"

// Client -
type Client struct {
	HostURL         string
	DatoEnvironment string
	HTTPClient      *http.Client
	Token           string
}

// NewClient -
func NewClient(host *string, authToken *string, datoEnvironment *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default DatoCMS URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if authToken != nil {
		c.Token = *authToken
	}

	if datoEnvironment != nil {
		c.DatoEnvironment = *datoEnvironment
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	tokenParam := "Bearer " + c.Token
	// Set required headers
	req.Header.Set("accept", "application/json") 
	req.Header.Set("content-type", "application/json")

	req.Header.Set("Authorization", tokenParam) // Set the Auth Header
	req.Header.Set("X-Environment", c.DatoEnvironment) // Set the Dato Environment Header
	req.Header.Set("X-Api-Version", "3") // Set the Dato Environment Header


	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// fmt.Println(req.Response.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusCreated {
		fmt.Println(res.StatusCode)
		return body, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}