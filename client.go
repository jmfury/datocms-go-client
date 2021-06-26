package datocms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default DatoCMS URL
const HostURL string = "https://site-api.datocms.com/site"

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

	req.Header.Set("Authorization", tokenParam)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
