package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateModel
func (c *Client) CreateModel(modelParam Model) (*Model, error) {
	dto := ModelDto{}
	dto.Data = modelParam
	rb, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/item-types", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	response := ModelDto{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	newModel := response.Data


	return &newModel, nil
}


// GetModel
func (c *Client) GetModel(apiKey string) (*Model, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/item-types/%s", c.HostURL, apiKey), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := ModelDto{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	responseModel := response.Data


	return &responseModel, nil
}

// UpdateModel 
func (c *Client) UpdateModel(modelParam Model) (*Model, error) {
	dto := ModelDto{}
	dto.Data = modelParam
	rb, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/item-types/%s", c.HostURL, modelParam.Attributes.ApiKey), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := ModelDto{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	updatedModel := response.Data


	return &updatedModel, nil
}

// DeleteModel
func (c *Client) DeleteModel(apiKey string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/item-types/%s", c.HostURL, apiKey), nil)
	if err != nil {
		return  err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return  err
	}

	response := DeleteModelResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return  err
	}
	return nil
}