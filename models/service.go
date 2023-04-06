package models

import "net/http"

type Service struct {
	Key       string
	Name      string
	Version   string
	Host      string
	Port      string
	Namespace string
}

type HTTPRoute struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type ApiResponse struct {
	Status  int         `json:"status"`
	Title   string      `json:"title"`
	Detail  interface{} `json:"detail"`
	Success bool        `json:"success"`
}
