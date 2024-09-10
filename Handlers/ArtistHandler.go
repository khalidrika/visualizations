package groupie

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"

	groupie "groupie/data"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) { // traiter les information des artistes dans la second page
	url1 := "https://groupietrackers.herokuapp.com/api/"
	var data groupie.Artist
	var wg sync.WaitGroup // un variable waitgroup pour gerer les gourotines

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	id := strings.Trim(r.URL.Path, "/artist/")
	if len(id) > 500 {
		ErrorHandler(w, r, http.StatusNotFound, "page Not found")
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil || num <= 0 || num > 52 {
		ErrorHandler(w, r, http.StatusNotFound, "page Not found")

		return
	}
	wg.Add(4) // declarer a waitgroup que ona 4 gourotine

	go func() {
		defer wg.Done() // lorsque la gourotine terminer envouyer un Done

		FetchHandler(url1+"locations/", &data.Location, strconv.Itoa(num), w, r)
	}()
	go func() {
		defer wg.Done()
		FetchHandler(url1+"dates/", &data.Dates, strconv.Itoa(num), w, r) // remplir les defèrants structures à partir des donnés des APIS
	}()
	go func() {
		defer wg.Done()

		FetchHandler(url1+"artists/", &data.Information, strconv.Itoa(num), w, r)
	}()
	go func() {
		defer wg.Done()
		FetchHandler(url1+"relation/", &data.Rolation, strconv.Itoa(num), w, r)
	}()
	wg.Wait() // attendre l'exucution de touts les gourotine avant de continue l'execution de programme
	
	data.Dates.Dates = RemoveAsterisk(data.Dates.Dates)

	tmpl, err2 := template.ParseFiles("templete/artist.html")
	if err2 != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return

	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
