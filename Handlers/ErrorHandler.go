package groupie

import (
	"net/http"
	"text/template"
)

type err struct {
	Error string
	Code  int
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, Message string) {
	w.WriteHeader(code)
	msg_error := &err{Error: Message + "!", Code: code}

	tmp, err := template.ParseFiles("templete/error.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, msg_error)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
