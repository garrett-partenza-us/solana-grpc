package main

import (
	"log"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type getAccountBalanceRequest struct {
	Account string `json:account`
}

type getAccountBalanceResponse struct {
	Slot string `json:"slot"`
	Value string  `json:"value"`
}

type getSlotLeaderRequest struct {
	Commitment string `json:"commitment"`
	MinContextSlot int64 `json:"minContextSlot"`
}

type getSlotLeaderResponse struct {
	Pubkey string `json:"pubkey"`
}

func main() {
	wsURL := "ws://localhost:8080/v1/getSlotLeaderStream"
	runGetSlotLeader(wsURL)
}

func runGetSlotLeader(wsURL string) {
	reqBody := getSlotLeaderRequest {
		Commitment: "None",
		MinContextSlot: 0,
	}

	req, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Error marshalling getSlotLeaderRequest: ", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("Error dialing websocket: ", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, req)
	if err != nil {
		log.Fatal("Error sending request to websocket: ", err)
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Error reading message from websocket: ", err)
	}

	var response getSlotLeaderResponse
	err = json.Unmarshal(msg, &response)
	if err != nil {
		log.Fatal("Error unmarshalling websocket response: ", err)
	}

	log.Println("Slot leader: ", response.Pubkey)

}
