package pokeapi

import (
	"net/http"
	"time"

	"github.com/amayakmt/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{},
		cache:      *pokecache.NewCache(interval),
	}
}
