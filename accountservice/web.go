package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type Route struct {
	Name       string
	Method     string
	Patten     string
	HandleFunc http.HandlerFunc
}

var Routes = []Route{
	{
		Name:       "GetUser",
		Method:     "GET",
		Patten:     "/accounts/{accountId}",
		HandleFunc: GetUser,
	},
}

func NewHHTTPRouter() *mux.Router {
	var router = new(mux.Router)
	for _, r := range Routes {
		router.Path(r.Patten).Methods(r.Method).Name(r.Name).Handler(r.HandleFunc)
	}
	return router
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	accountId := mux.Vars(r)["accountId"]

	log.Infof("get account of %s", accountId)

	var ret = struct {
		AccountId string `json:"accountId"`
	}{AccountId: accountId}

	body, _ := json.Marshal(ret)

	WriteHTTPJSON(body, w)
}

func WriteHTTPJSON(body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-length", strconv.Itoa(len(body)))

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
