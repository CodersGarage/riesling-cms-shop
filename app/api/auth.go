package api

import (
	"net/http"
	"github.com/dghubble/sling"
	"github.com/spf13/viper"
	"encoding/json"
	"riesling-cms-shop/app/utils"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type AuthResponse struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
}

const (
	ACCESS_TOKEN = "access_token"
	HASH         = "hash"
)

func CheckAuth(accessToken string, hash string) *AuthResponse {
	req, err := sling.New().Get(viper.GetString("authorization.url")).Set(HASH, hash).Set(ACCESS_TOKEN, accessToken).Request()
	if err != nil {
		utils.LogP("Auth", err)
		return nil
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogP("Auth", err)
		return nil
	}
	authResponse := &AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		utils.LogP("Auth", err)
		return nil
	}
	return authResponse
}

func AdminAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get(HASH)
		accessToken := r.Header.Get(ACCESS_TOKEN)
		resp := CheckAuth(accessToken, hash)
		if resp != nil && resp.Code == 200 && resp.Data["is_valid"].(bool) && resp.Data["level"].(int) == 1 {
			h.ServeHTTP(w, r)
			return
		}
		res := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(res, w)
	}
}

func SelfAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get(HASH)
		accessToken := r.Header.Get(ACCESS_TOKEN)
		resp := CheckAuth(accessToken, hash)
		if resp != nil && resp.Code == 200 && resp.Data["is_valid"].(bool) {
			h.ServeHTTP(w, r)
			return
		}
		res := APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized request",
		}
		ServeAsJSON(res, w)
	}
}
