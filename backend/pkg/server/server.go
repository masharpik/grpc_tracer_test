package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/test", Test).Methods("GET")
	return r
}

func Run(handler http.Handler, host string, port string) error {
	log.Println("Backend server starts")
	return http.ListenAndServe(host+":"+port, handler)
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("HI")
}
