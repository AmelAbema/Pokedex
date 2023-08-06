package pokeapi

import (
	"github.com/AmelAbema/Pokedex/internal/pokecache"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {

	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
