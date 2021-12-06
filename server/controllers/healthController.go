package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Health godoc
// @Summary      Health Check
// @Description  get server health
// @Tags         health
// @Produce      application/json
// @Success      200  {array}   Result
// @Failure      500
// @Router       /health [get]
func (c *ControllerManager) HealthCheck(resp http.ResponseWriter, req *http.Request) {
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
