package handler

import (
	"encoding/json"
	"fmt"
	"github.com/wisp-gg/gamequery"
	"github.com/wisp-gg/gamequery/api"
	"net/http"
	"os"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort, _ := strconv.ParseUint(os.Getenv("SERVER_PORT"), 10, 16)

	res, err := gamequery.Query(api.Request{
		IP:   serverAddress,
		Port: uint16(serverPort),
		Game: "source",
	})

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	stats, err := json.Marshal(res)
	fmt.Fprintf(w, "%s", stats)
}
