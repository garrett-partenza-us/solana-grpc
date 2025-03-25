package methods

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type getAccountBalanceJSONRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			APIVersion string `json:"apiVersion"`
			Slot       int64  `json:"slot"`
		} `json:"context"`
		Value int64 `json:"value"`
	} `json:"result"`
	ID int64 `json:"id"`
}

func GetAccountBalanceJSONRPC(account string) getAccountBalanceJSONRPCResponse {
	request := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "getBalance",
		Params:  []interface{}{account},
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Failed to marshal getAccountBalanceJSONRPC request", err)
	}

	resp, err := http.Post("http://nuc:8899", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal("Failed to send getAccountBalanceJSONRPC request", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read getAccountBalanceJSONRPC response body: ", err)
	}

	var jsonResponse getAccountBalanceJSONRPCResponse
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		log.Fatal("Failed to unmarshal getAccountBalanceJSONRPCResponse to object: ", err)
	}

	return jsonResponse

}
