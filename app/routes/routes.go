package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/spf13/viper"
	"sync"
)

var RootRoute = mux.NewRouter()
var WaitGroup = sync.WaitGroup{}

func InitRoutes() {
	go http.ListenAndServe(viper.GetString("app.uri"), RootRoute)
	WaitGroup.Add(1)
}
