package handler

import "net/http"

func HealthCheck() (string, func(resp http.ResponseWriter, req *http.Request)) {
	return "/health", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(http.StatusOK)
	}
}
