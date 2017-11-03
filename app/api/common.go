package api

import (
	"encoding/json"
	"net/http"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ResponseValue map[string]interface{}

func ServeAsJSON(resp APIResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

func GetURLParam(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}
