package musicbrainz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

const (
	userAgentString = "Apollo/2.0.0 (i.acnaylor@gmail.com)"
	baseUrl         = "https://musicbrainz.org/ws/2"
)

func createRequest(endpoint string) (*http.Request, error) {
	request, err := http.NewRequest("GET", endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", userAgentString)

	return request, nil
}

type Client struct {
	limiter rate.Limiter
}

func NewClient() Client {
	return Client{
		limiter: *rate.NewLimiter(rate.Every(time.Second), 1),
	}
}

func (c *Client) SearchArtistByName(artistName string) (*SearchResult, error) {
	if !c.limiter.Allow() {
		c.limiter.Wait(context.Background())
	}

	endpoint := fmt.Sprintf(
		"%s/artist?query=artist:%s&limit=3",
		baseUrl, url.QueryEscape(artistName),
	)

	request, err := createRequest(endpoint)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status indicated error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(body, &result)

	return &result, err
}

func (c *Client) LookupArtistById(artistMbid string) (*Artist, error) {
	if !c.limiter.Allow() {
		c.limiter.Wait(context.Background())
	}

	endpoint := fmt.Sprintf(
		"%s/artist/%s?inc=release-groups+genres",
		baseUrl, artistMbid,
	)

	request, err := createRequest(endpoint)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status indicated error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result Artist
	err = json.Unmarshal(body, &result)

	return &result, err
}
