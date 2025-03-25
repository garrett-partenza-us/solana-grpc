package methods

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type getLatestBlockHashJSONRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			APIVersion string `json:"apiVersion"`
			Slot       int64  `json:"slot"`
		} `json:"context"`
		Value struct {
			Blockhash            string `json:"blockhash"`
			LastValidBlockHeight int64  `json:"lastValidBlockHeight"`
		} `json:"value"`
	} `json:"result"`
	ID int64 `json:"id"`
}

func GetLatestBlockHashJSONRPC() getLatestBlockHashJSONRPCResponse {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "getLatestBlockhash",
		Params:  []interface{}{},
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Failed to marshal getLatestBlockHashJSONRPC request", err)
	}

	resp, err := http.Post("http://nuc:8899", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal("Failed to send getLatestBlockHashJSONRPC request", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read getLatestBlockHashJSONRPC response body: ", err)
	}

	var jsonResponse getLatestBlockHashJSONRPCResponse
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		log.Fatal("Failed to unmarshal getLatestBlockhashJSONRPCResponse to object: ", err)
	}

	return jsonResponse

}
