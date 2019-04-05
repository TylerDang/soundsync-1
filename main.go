package main

import (
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joshuaj1397/soundsync/api"
)

var (
	port         = "3005"
	mySigningKey = []byte("ASuperSecretSigningKeyCreatedByTheAliensFromArrival")
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	router := mux.NewRouter()

	// API
	router.Handle("/GetToken", api.GetToken).Methods("GET")
	router.Handle("/CreateParty/{nickname}/{phoneNum}/{partyName}", jwtMiddleware.Handler(api.CreateParty)).Methods("POST")
	router.Handle("/JoinParty/{nickname}/{partyCode}/{phoneNum}", jwtMiddleware.Handler(api.JoinParty)).Methods("POST")
	// router.HandleFunc("/Verify/{phoneNum}/{name}/{authCode}", api.Verify).Methods("POST")

	//TODO: Find out what this endpoint needs and returns
	router.HandleFunc("/LinkSpotify", api.LinkSpotify).Methods("GET")
	router.HandleFunc("/callback", api.SpotifyCallback).Methods("GET")
	router.HandleFunc("/MediaControls/Play", api.Play).Methods("PUT")
	router.HandleFunc("/MediaControls/Pause", api.PlayPause).Methods("PUT")
	router.HandleFunc("/MediaControls/Next", api.NextPrev).Methods("POST")
	router.HandleFunc("/MediaControls/Previous", api.NextPrev).Methods("POST")
	router.HandleFunc("/SearchSpotify/{query}", api.SearchSpotify).Methods("GET")
	router.HandleFunc("/AddSong/{songURI}", api.AddSong).Methods("POST")
	// router.HandleFunc("/SongQueue", api.SongQueue).Methods("GET")
	// router.HandleFunc("/CurrentSong", api.MediaControls).Methods("GET")
	// router.HandleFunc("/SkipSong/{songId}/{partyId}", api.SkipSong).Methods("POST")
	// router.HandleFunc("/RemoveSong/{songId}/{partyI 	d}", api.RemoveSong).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, router))
	fmt.Println("Listening on port: " + port)
}
