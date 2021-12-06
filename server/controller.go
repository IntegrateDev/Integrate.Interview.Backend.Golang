package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Result struct {
	Result interface{}   `json:"result"`
	Errors []interface{} `json:"errors"`
}

// Health godoc
// @Summary      Health Check
// @Description  get server health
// @Tags         health
// @Produce      json
// @Success      200  {array}   Result
// @Failure      500
// @Router       /health [get]
func healthCheck(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	responseObj := Result{}
	responseObj.Result = "Everything is peachy, server up"
	responseObj.Errors = make([]interface{}, 0)
	respData, err := json.Marshal(responseObj)
	if err != nil {
		log.Printf("failed to marshal healthcheck with: %s\n", err)
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(200)
	resp.Write(respData)
}
