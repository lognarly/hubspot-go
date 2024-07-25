package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	DefaultAddress = "https://api.hubapi.com"
)

var (
	ErrMissingToken = "an API token must be provided to call the Hubspot API"
)

type Client struct {
	baseURL string
	token   string
	http    *http.Client

	Associations        Associations
	Calls               Calls
	Companies           Companies
	Contacts            Contacts
	Deals               Deals
	Emails              Emails
	FeedbackSubmissions FeedbackSubmissions
	LineItems           LineItems
	Meetings            Meetings
	Notes               Notes
	Owners              Owners
	Pipelines           Pipelines
	Products            Products
	Tasks               Tasks
	Tickets             Tickets
	Quotes              Quotes
}

// NewHubspotClient Used to create a new HubSpot Client
func NewHubspotClient(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf(ErrMissingToken)
	}
	return newHubspotClientWithDefaults(token), nil
}

// NewHubspotClientFromHttpClient Creates a new HubSpot Client, but allows for passing in a custom HTTP client.
// This can be used for passing contexts throughout SDK usage for additional customization.
func NewHubspotClientFromHttpClient(token string, httpClient *http.Client) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf(ErrMissingToken)
	}
	client := newHubspotClientWithDefaults(token)
	if httpClient != nil {
		client.http = httpClient
	}
	return client, nil
}

// Creates a new HubSpot client with defaults
func newHubspotClientWithDefaults(token string) *Client {
	client := &Client{
		baseURL: DefaultAddress,
		token:   token,
		http: &http.Client{
			Timeout: time.Duration(30 * time.Second),
		},
	}
	client.Associations = &associations{client: client}
	client.Calls = &calls{client: client}
	client.Companies = &companies{client: client}
	client.Contacts = &contacts{client: client}
	client.Deals = &deals{client: client}
	client.Emails = &emails{client: client}
	client.FeedbackSubmissions = &feedbackSubmissions{client: client}
	client.LineItems = &lineItems{client: client}
	client.Meetings = &meetings{client: client}
	client.Notes = &notes{client: client}
	client.Owners = &owners{client: client}
	client.Pipelines = &pipelines{client: client}
	client.Products = &products{client: client}
	client.Tasks = &tasks{client: client}
	client.Tickets = &tickets{client: client}
	client.Quotes = &quotes{client: client}

	return client
}

func (c *Client) newHttpRequest(ctx context.Context, method string, endpoint string, v interface{}) (*http.Request, error) {
	var err error
	var body []byte
	var newBody io.Reader
	u, err := c.formatUrl(endpoint)
	if err != nil {
		return nil, err
	}

	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	switch method {
	case "GET", "DELETE":
		if v != nil {
			q, err := query.Values(v)
			if err != nil {
				return nil, err
			}
			u.RawQuery = u.RawQuery + encodeQueryParams(q)
		}
	case "POST", "PUT", "PATCH":
		if v != nil {
			if body, err = json.Marshal(v); err != nil {
				return nil, err
			}
			newBody = bytes.NewReader(body)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), newBody)
	if err != nil {
		return nil, err
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	statusOk := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOk {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		errMsg := fmt.Sprintf("%d: %s", res.StatusCode, string(resBody))
		return fmt.Errorf(errMsg)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(resBody) > 0 {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) formatUrl(endpoint string) (*url.URL, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	u.RawQuery = q.Encode()

	return u, nil
}

func encodeQueryParams(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) > 1 {
			val := strings.Join(vs, ",")
			vs = vs[:0]
			vs = append(vs, val)
		}
		keyEscaped := url.QueryEscape(k)

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
