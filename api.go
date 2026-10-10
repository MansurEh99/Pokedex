package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MansurEh99/pokedexcli/internal/pokecache"
)

func fetchData(url string, cache *pokecache.Cache) ([]byte, error) {
	if data, ok := cache.Get(url); ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, data)

	return data, nil
}
