package RadishCMS

import (
    "net/http"
    "github.com/gorilla/mux"
)

func init() {
    r := mux.NewRouter()
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", r)
}