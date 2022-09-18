package rest_client

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	Address   string        `yaml:"address"`
	Timeout   time.Duration `yaml:"timeout"`
	DebugMode bool          `yaml:"debug_mode"`
}

type client struct {
	options Options
	client  *resty.Client
}

func NewRestClient(options Options) RestClient {
	httpClient := resty.New()

	httpClient.SetTimeout(options.Timeout)
	httpClient.SetDebug(options.DebugMode)

	return &client{
		options: options,
		client:  httpClient,
	}
}

type RestClient interface {
	Get(path string, header http.Header, payload interface{}) (body []byte, httpStatusCode int, err error)
}

func (c *client) Get(path string, header http.Header, payload interface{}) (body []byte, httpStatusCode int, err error) {
	url := c.options.Address + path

	request := c.client.R()
	request.Header = header

	if header["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}

	httpResp, httpErr := request.Get(url)
	if httpErr != nil {
		return
	}

	httpStatusCode = httpResp.StatusCode()
	body = httpResp.Body()

	return
}
