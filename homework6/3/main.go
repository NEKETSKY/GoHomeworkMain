package main

import (
	"fmt"
	"log"
	"net/http"
)

func postGet(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 page not found.", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "template.html")
	}

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		address := r.FormValue("address")
		token := fmt.Sprintf("%v:%v", name, address)

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
		})

		http.ServeFile(w, r, "template.html")
	}
}

func main() {

	fmt.Printf("Server starting on port: 8080")

	http.HandleFunc("/", postGet)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
