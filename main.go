package main

import (
	"time"

	"github.com/navacodes/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Minute),
	}
	startRepl(cfg)

}
