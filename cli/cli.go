package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/piann/coin_101/explorer"
	"github.com/piann/coin_101/rest"
)

func usage() {
	fmt.Println("Usage )")
	fmt.Println("-port : Choose port number")
	fmt.Println("-mode : Choose mode between REST 'rest' and 'html'")
	os.Exit(0)
}

func Start() {

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose mode between REST 'rest' and 'html'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

}
