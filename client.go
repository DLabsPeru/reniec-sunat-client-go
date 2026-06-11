package reniecsunatclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    http.Header
}

type Option func(*Client)

const DefaultBaseURL = "https://api-reniec-sunat.destiny-peru.com"

func New(opts ...Option) *Client {
	return NewWithBaseURL(DefaultBaseURL, opts...)
}

func NewWithBaseURL(baseURL string, opts ...Option) *Client {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(http.Header),
	}

	client.headers.Set("Accept", "application/json")
	client.headers.Set("User-Agent", "reniec-sunat-client/1.0")

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(client *Client) {
		if httpClient != nil {
			client.httpClient = httpClient
		}
	}
}

func WithHeader(key, value string) Option {
	return func(client *Client) {
		if strings.TrimSpace(key) != "" {
			client.headers.Set(key, value)
		}
	}
}

func (c *Client) GetPersonByDNI(ctx context.Context, dni string) (*Person, error) {
	var envelope Response[Person]
	if err := c.get(ctx, "/api/v1/persons/"+strings.TrimSpace(dni), &envelope); err != nil {
		return nil, err
	}

	return &envelope.Data, nil
}

func (c *Client) GetCompanyByRUC(ctx context.Context, ruc string) (*Company, error) {
	var envelope Response[Company]
	if err := c.get(ctx, "/api/v1/companies/"+strings.TrimSpace(ruc), &envelope); err != nil {
		return nil, err
	}

	return &envelope.Data, nil
}

func (c *Client) get(ctx context.Context, path string, target any) error {
	if c == nil {
		return fmt.Errorf("nil client")
	}
	if c.baseURL == "" {
		return fmt.Errorf("empty base URL")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return err
	}

	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, target); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	switch typed := target.(type) {
	case *Response[Person]:
		if resp.StatusCode >= 400 {
			return apiErrorFromEnvelope(resp.StatusCode, typed)
		}
	case *Response[Company]:
		if resp.StatusCode >= 400 {
			return apiErrorFromEnvelope(resp.StatusCode, typed)
		}
	}

	if resp.StatusCode >= 400 {
		return &Error{
			StatusCode: resp.StatusCode,
			Message:    http.StatusText(resp.StatusCode),
		}
	}

	return nil
}

func apiErrorFromEnvelope[T any](statusCode int, envelope *Response[T]) error {
	if envelope == nil {
		return &Error{StatusCode: statusCode, Message: http.StatusText(statusCode)}
	}

	errValue := &Error{
		StatusCode: statusCode,
		Message:    envelope.Message,
		RequestID:  envelope.RequestID,
		Path:       envelope.Path,
	}

	if envelope.Error != nil {
		errValue.Code = envelope.Error.Code
		if strings.TrimSpace(envelope.Error.Message) != "" {
			errValue.Message = envelope.Error.Message
		}
	}

	if strings.TrimSpace(errValue.Message) == "" {
		errValue.Message = http.StatusText(statusCode)
	}

	return errValue
}
