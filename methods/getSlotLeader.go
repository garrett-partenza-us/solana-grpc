package methods

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type getSlotLeaderJSONRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result string `json:"result"`
	ID int64 `json:"id"`
}

func GetSlotLeaderJSONRPC() getSlotLeaderJSONRPCResponse {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "getSlotLeader",
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Failed to marshal getSlotLeaderJSONRPC request", err)
	}

	resp, err := http.Post("http://nuc:8899", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal("Failed to send getSlotLeaderJSONRPC request", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read getSlotLeaderJSONRPC response body: ", err)
	}

	var jsonResponse getSlotLeaderJSONRPCResponse
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		log.Fatal("Failed to unmarshal getSlotLeaderJSONRPCResponse to object: ", err)
	}

	return jsonResponse

}
