package groupie

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"
)

// func StyleHandler(w http.ResponseWriter, r *http.Request) {
// 	if strings.HasSuffix(r.URL.Path, "/") {
// 		ErrorHandler(w, r, http.StatusNotFound, "page not found")
// 		return
// 	}
// 	// Serve static files from the "./style/" directory
// 	fs := http.FileServer(http.Dir("./style/"))
// 	http.StripPrefix("/style/", (fs)).ServeHTTP(w, r)
// }

// Handle serving css content, while blocking access to paths "/css/..."
func StyleHandler(w http.ResponseWriter, r *http.Request) {
	// Strip the "/css/" prefix from the URL path to get the relative file path
	a := "style/styles/" + strings.TrimPrefix(r.URL.Path, "/style/")
	// Read the file from the embedded filesystem

	b, err := os.ReadFile(a)
	if err != nil {
		ErrorHandler(w, r, 404, "page not fond")
		return
	}
	// Serve the file content
	http.ServeContent(w, r, a, time.Now(), bytes.NewReader(b))
}
