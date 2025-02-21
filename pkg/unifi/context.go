package unifi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Context struct {
	ApiKey   string
	PageSize int
	Client   *http.Client
}

type Status struct {
	Error          error
	Code           string `json:"code"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	Message        string `json:"message"`
	TraceID        string `json:"traceId"`
	NextToken      string `json:"nextToken"`
}

func (u *Context) Get(uri string, params map[string]string, resp any) Status {
	// build URL with host, uri, and params.  Substitue params uri if specified.
	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Add(k, v)
		}
	}

	// Build full URL
	fullURL := fmt.Sprintf("https://api.ui.com%s", uri)
	if len(values) > 0 {
		fullURL = fmt.Sprintf("%s?%s", fullURL, values.Encode())
	}

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Status{Error: fmt.Errorf("creating request: %w", err)}
	}

	// Add headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-KEY", u.ApiKey)

	// Make request
	response, err := u.Client.Do(req)
	if err != nil {
		return Status{
			Error: fmt.Errorf("request: %s: %w", fullURL, err),
		}
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Status{Error: fmt.Errorf("reading response: %w", err)}
	}

	// Check status code
	if response.StatusCode != http.StatusOK {
		status := Status{}
		if err := json.Unmarshal(body, &status); err != nil {
			return Status{
				Error:          fmt.Errorf("unexpected status code: %d", response.StatusCode),
				HTTPStatusCode: response.StatusCode,
			}
		}
		status.HTTPStatusCode = response.StatusCode
		return status
	}

	// Unmarshal response
	if err := json.Unmarshal(body, resp); err != nil {
		return Status{Error: fmt.Errorf("unmarshaling response: %w", err)}
	}

	return Status{
		HTTPStatusCode: response.StatusCode,
	}
}

func (u *Context) Post(uri string, params map[string]string, body any, resp any) Status {
	return Status{}
}

func (s *Status) Failed() bool {
	return s.Error != nil
}
