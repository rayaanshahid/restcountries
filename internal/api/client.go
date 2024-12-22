package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rayaanshahid/restcountries/internal/models"
	"golang.org/x/net/http2"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) FetchCountries() ([]models.Country, error) {
	transport := &http2.Transport{
		StrictMaxConcurrentStreams: true,
	}

	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(fmt.Sprintf("%s/all", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var countries []models.Country
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}

func (c *Client) FetchEuropeanCountries() ([]models.Country, error) {
	transport := &http2.Transport{
		StrictMaxConcurrentStreams: true,
	}

	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(fmt.Sprintf("%s/region/europe", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var countries []models.Country
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}
