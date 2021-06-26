package datocms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetFields() (*FieldsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fields", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	fields := FieldsResponse{}
	err = json.Unmarshal(body, &fields)
	if err != nil {
		return nil, err
	}

	return &fields, nil
}

func (c *Client) UpdateFieldAttributes(apiKey string, attributes Attributes) (*FieldResponse, error) {
	rb, err := json.Marshal(attributes)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/fields/%s", c.HostURL, apiKey), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	field := FieldResponse{}
	err = json.Unmarshal(body, &field)
	if err != nil {
		return nil, err
	}

	return &field, nil
}
