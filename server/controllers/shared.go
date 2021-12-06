package controllers

import "net/http"

type Result struct {
	Result interface{}   `json:"result"`
	Errors []interface{} `json:"errors"`
}

type Lead struct {
}

type ControllerManager struct {
}

type Manager interface {
	AllLeads(resp http.ResponseWriter, req *http.Request)
	HealthCheck(resp http.ResponseWriter, req *http.Request)
}

func NewController(controller *ControllerManager) Manager {
	return controller
}
