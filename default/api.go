package RadishCMS

import (
  // "google.golang.org/appengine/datastore"
  "net/http"
  "encoding/json"
)

type GetSettingUpT bool

func GetSettingUp(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(GetSettingUpT(true))
}
