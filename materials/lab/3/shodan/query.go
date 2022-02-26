package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	queryTagsPath   = "/shodan/query/tags"
	querySearchPath = "/shodan/query/search"
	queryPath       = "/shodan/query"
)

// QueryTagsMatch represents a matched tag.
type QueryTagsMatch struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// QueryTags represents matched tags.
type QueryTags struct {
	Total   int               `json:"total"`
	Matches []*QueryTagsMatch `json:"matches"`
}

// QuerySearchMatch is a match of QuerySearch.
type QuerySearchMatch struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Query       string   `json:"query"`
	Votes       int      `json:"votes"`
	Timestamp   string   `json:"timestamp"`
	Tags        []string `json:"tags"`
}

// QuerySearch is the results of querying saved search queries.
type QuerySearch struct {
	Total   int                 `json:"total"`
	Matches []*QuerySearchMatch `json:"matches"`
}

// GetQueryTags obtains a list of popular tags for the saved search queries in Shodan.
func (c *Client) GetQueryTags(size string) (*QueryTags, error) {

	//req, err := c.NewRequest("GET", queryTagsPath, options, nil)
	res, err := http.Get(
		fmt.Sprintf("%s%s?key=%s&size=%s", BaseURL, queryTagsPath, c.apiKey, size),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret QueryTags
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// GetQueries obtains a list of search queries that users have saved in Shodan.
func (c *Client) GetQueries(p string, s string, o string) (*QuerySearch, error) {

	res, err := http.Get(
		fmt.Sprintf("%s%s?key=%s&size=%s&sort=%s&order=%s", BaseURL, queryPath, c.apiKey, p, s, o),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret QuerySearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil

}

// SearchQueries searches the directory of search queries that users have saved in Shodan.
func (c *Client) SearchQueries(q string, p string, s string, o string) (*QuerySearch, error) {

	if q == "" {
		return nil, fmt.Errorf("Please add query argument for search query")
	}
	res, err := http.Get(
		fmt.Sprintf("%s%s?key=%s&query=%s&size=%s&sort=%s&order=%s", BaseURL, querySearchPath, c.apiKey, q, p, s, o),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret QuerySearch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil

}