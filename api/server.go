package handler

import (
	"encoding/json"
	"fmt"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/wisp-gg/gamequery"
	"github.com/wisp-gg/gamequery/api"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ServerStateResponse struct {
	Name    string `json:"Name"`
	Players struct {
		Current int         `json:"Current"`
		Max     int         `json:"Max"`
		Names   interface{} `json:"Names"`
	} `json:"Players"`
	Raw struct {
		Protocol    int    `json:"Protocol"`
		Name        string `json:"Name"`
		Map         string `json:"Map"`
		Folder      string `json:"Folder"`
		Game        string `json:"Game"`
		ID          int    `json:"ID"`
		Players     int    `json:"Players"`
		Maxplayers  int    `json:"MaxPlayers"`
		Bots        int    `json:"Bots"`
		Servertype  int    `json:"ServerType"`
		Environment int    `json:"Environment"`
		Visibility  int    `json:"Visibility"`
		Vac         int    `json:"VAC"`
		Version     string `json:"Version"`
		Edf         int    `json:"EDF"`
		Extradata   struct {
			Port         int    `json:"Port"`
			Steamid      int64  `json:"SteamID"`
			Sourcetvport int    `json:"SourceTVPort"`
			Sourcetvname string `json:"SourceTVName"`
			Keywords     string `json:"Keywords"`
			Gameid       int    `json:"GameID"`
		} `json:"ExtraData"`
	} `json:"Raw"`
}

type ServerState struct {
	ID          string `fauna:"id"`
	Name        string `fauna:"name"`
	PlayerCount int    `fauna:"playerCount"`
	Version     string `fauna:"version"`
	Timestamp   int64  `fauna:"timestamp"`
}

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

	serverStateJson, err := json.Marshal(res)

	query := r.URL.Query()
	persist := query.Get("persist")
	if persist == "y" {
		persistState(serverStateJson)
	}

	fmt.Fprintf(w, "%s", serverStateJson)
}

func persistState(serverStateJson []byte) {
	client := f.NewFaunaClient(os.Getenv("FAUNADB_SERVER_KEY"))

	var serverStateResponse ServerStateResponse
	err := json.Unmarshal(serverStateJson, &serverStateResponse)
	if err != nil {
		log.Println(err)
	}

	var serverState = ServerState{
		ID:          strconv.Itoa(serverStateResponse.Raw.ID),
		Name:        serverStateResponse.Name,
		PlayerCount: serverStateResponse.Players.Current,
		Version:     serverStateResponse.Raw.Extradata.Keywords,
		Timestamp:   time.Now().Unix(),
	}

	_, err = client.Query(f.Create(f.Collection("ServerState"), f.Obj{"data": serverState}))
	if err != nil {
		log.Println(err)
	}
}
