package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/spf13/viper"
	"sync"
	"riesling-cms-shop/app/api"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var RootRoute = mux.NewRouter()
var WaitGroup = sync.WaitGroup{}

func InitRoutes() {
	v1 := RootRoute.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/product", api.SelfAuth(api.CreateProduct)).Methods("POST")
	v1.HandleFunc("/product", api.SelfAuth(api.FindProducts)).Methods("GET")

	go http.ListenAndServe(viper.GetString("app.uri"), RootRoute)
	WaitGroup.Add(1)
}
