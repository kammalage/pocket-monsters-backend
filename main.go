package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mtslzr/pokeapi-go"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pokedex", PokedexHandler).Methods(http.MethodGet)
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := []byte("<div><h1>Home Page</h1></div>")

	w.Write(html)
}

func PokedexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pokedex, err := pokeapi.Pokedex("kanto")
	if err != nil {
		log.Fatal(err)
	}

	jsonPokedex, err := json.Marshal(pokedex.PokemonEntries)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonPokedex)
}
