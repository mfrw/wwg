package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		defer trackTime(time.Now(), "TIME")
		log.Printf("Logged connection from %s to %s\n", r.RemoteAddr, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// START OMIT
func main() {
	router := mux.NewRouter()
	router.Use(simpleMw)

	router.HandleFunc("/", IndexPage).Methods("GET")
	router.HandleFunc("/chart", chart).Methods("GET")
	router.HandleFunc("/data", data).Methods("GET")
	router.HandleFunc("/graph", graph).Methods("GET")
	router.HandleFunc("/style", style).Methods("GET")

	fmt.Println("Starting Server on : [localhost:8080]")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// END OMIT

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "final.html")
}

func chart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "js/chart.js")
}
func style(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "css/style.css")
}

func data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"delayedPrice\":%f, \"delayedPriceTime\":%d }", rand.Float32()*50.0, time.Now().Second())
}

func graph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, "js/plotter.js")
}

func data_web(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://api.iextrading.com/1.0/stock/aapl/delayed-quote")
	if err != nil {
		log.Println("Could not get latest data", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
