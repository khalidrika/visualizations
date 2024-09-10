package groupie

import (
	"encoding/json"
	"net/http"
)

func FetchHandler(url string, data interface{}, id string, w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(url + id)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(data)
}
