package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Health godoc
// @Summary      Get All Lead Data
// @Description  Returns Array of Lead Results
// @Tags         leads
// @Produce      application/json
// @Success      200  {array}   Result
// @Failure      500
// @Router       /leads [get]
func (c *ControllerManager) AllLeads(resp http.ResponseWriter, req *http.Request) {
	responseObj := Result{}
	responseObj.Result = make([]Lead, 0)
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
