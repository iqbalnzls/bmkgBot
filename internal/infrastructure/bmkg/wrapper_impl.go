package bmkg

import (
	"errors"
	"net/http"

	"github.com/bmkgBot/internal/shared/config"
	restClient "github.com/bmkgBot/internal/shared/rest_client"
)

type bmkgWrapper struct {
	client     restClient.RestClient
	bmkgConfig *config.BMKGConfig
}

func NewBMKGWrapperIFace(restClient restClient.RestClient, bmkgConfig *config.BMKGConfig) BMKGWrapperIFace {
	if restClient == nil {
		panic("rest client is nil")
	}
	if bmkgConfig == nil {
		panic("bmkg config is nil")
	}

	return &bmkgWrapper{
		client:     restClient,
		bmkgConfig: bmkgConfig,
	}
}

func (w *bmkgWrapper) Get(key string) (body []byte, err error) {
	path, ok := w.bmkgConfig.Path[key]
	if !ok {
		err = errors.New("path not found")
		return
	}

	body, status, err := w.client.Get(path, http.Header{}, nil)
	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = errors.New("response with no success")
		return
	}

	if body == nil {
		err = errors.New("body is nil")
		return
	}

	return
}
