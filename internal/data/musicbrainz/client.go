package musicbrainz

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	userAgentString = "Apollo/1.0.0 (i.acnaylor@gmail.com)"
	baseUrl         = "https://musicbrainz.org/ws/2"
)

func createRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", userAgentString)

	return request, nil
}

// Perform a search based on the artists name, selecting the top 3 results
func SearchArtist(artistName string) (*Search, error) {
	url := fmt.Sprintf(
		"%s/artist?query=artist:%s&limit=3",
		baseUrl, url.QueryEscape(artistName),
	)

	request, err := createRequest(url)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response was not ok: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result Search
	err = json.Unmarshal(body, &result)

	return &result, err
}

// Using the artists MusicBrainz ID, find information about them
func LookupArtist(artistMbid string) (*Artist, error) {
	url := fmt.Sprintf(
		"%s/artist/%s?inc=release-groups+genres",
		baseUrl, artistMbid,
	)

	request, err := createRequest(url)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response was not ok: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result Artist
	err = json.Unmarshal(body, &result)

	return &result, err
}
