package main

import (
	"github.com/piann/coin_101/explorer"
	"github.com/piann/coin_101/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
