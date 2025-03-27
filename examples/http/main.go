package main

import (
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
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

	runGetLatestBlockHash()
	runGetAccountBalance()
	runGetSlotLeader()
}

func runGetLatestBlockHash() {
	resp, err := http.Get("http://localhost:8080/v1/getLatestBlockHash")
	if err != nil {
		log.Fatal("Error calling getLatestBlockHash endpoint: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Failed to read response body: ", err)
		}
		log.Println("Response body: ", string(body))
	} else {
		log.Fatal("Failed to read getLatestBlockHash response body: ", err)
	}
}

func runGetAccountBalance() {
	reqBody := getAccountBalanceRequest{
		Account: "6mqzvkBvdfSVhmCts4iLyBCxySigk8RMD3Uv9HVyPqQ8",
	}

	req, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Failed to marshal request for getAccountBalance: ", err)
	}

	resp, err := http.Post("http://localhost:8080/v1/getAccountBalance", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Fatal("Failed to send post request to getAccountBalance: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var response getAccountBalanceResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Fatal("Error decoding response for getAccountBalance: ", err)
		}

		log.Println("Account balance: ", response.Value)
		log.Println("Slot: ", response.Slot)
	} else {
		log.Fatal("Failed to call getAccountBalance endpoint: ", err)
	}
}

func runGetSlotLeader() {
	reqBody := getSlotLeaderRequest{
		Commitment: "None",
		MinContextSlot: 1,
	}

	req, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Failed to marshal request for getSlotLeader: ", err)
	}

	resp, err := http.Post("http://localhost:8080/v1/getSlotLeader", "application/json", bytes.NewBuffer(req))
	if err != nil {
		log.Fatal("Failed to send post request to getSlotLeader: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var response getSlotLeaderResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Fatal("Error decoding response for getSlotLeader: ", err)
		}

		log.Println("Slot leader public key: ", response.Pubkey)
	} else {
		log.Fatal("Failed to call getSlotLeader endpoint: ", err)
	}
}
