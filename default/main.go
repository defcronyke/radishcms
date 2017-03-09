package RadishCMS

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func init() {
    r := mux.NewRouter()
    r.HandleFunc("/setting-up", GetSettingUp).Methods("GET")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/dist")))
    http.Handle("/", cors.Default().Handler(r))
}
