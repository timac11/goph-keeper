package http

import (
	"strings"

	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	client resty.Client
}

// NewHTTPClient constructor
func NewHTTPClient(address string) *HTTPClient {
	client := resty.New()

	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	client.SetBaseURL(address)

	return &HTTPClient{client: *client}
}

func (client *HTTPClient) Login(login, password string) {}

func (client *HTTPClient) Register(login, password string) {}

func (client *HTTPClient) GetList(token string) {}

func (client *HTTPClient) GetSecret(token, id string) {}

func (client *HTTPClient) DeleteSecret(token, id string) {}

func (client *HTTPClient) UploadSecret(token string, data any) {}
