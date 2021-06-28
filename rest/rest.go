package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/piann/coin_101/blockchain"
	"github.com/piann/coin_101/utils"
)

type url string

var port string

type addBlockBody struct {
	Message string
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	url         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			url:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			url:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())

	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)

	}

}

func Start(portNum int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", portNum)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Rest : Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
