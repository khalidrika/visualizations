package groupie

import (
	"net/http"
	"text/template"

	groupie "groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) { // fonction pour traiter les informations necessaire dans la premiere page
	var tableau []groupie.Band
	url := "https://groupietrackers.herokuapp.com/api/artists"

	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, "page Not found")
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")

		return
	}

	data := HomeApi(url, tableau) // decodé les donnès jeson et les stocker dans le variable tableau
	tmpl, err := template.ParseFiles("templete/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
