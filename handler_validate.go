package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type chirp struct {
		Body string `json:"body"`
	}

	type returnPass struct {
		Cleaned_Body string `json:"cleaned_body"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	ch := chirp{}
	err := decoder.Decode((&ch))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(ch.Body) > 140 {
		respBody := returnFalse{
			Error: "Chirp is too long",
		}
		w.WriteHeader(400)

		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
		}
		w.Write(dat)

		return
	} else {

		w.WriteHeader(200)

		respBody := returnPass{
			Cleaned_Body: checkProfane(ch.Body),
		}

		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
		}
		w.Write(dat)

	}
}

func checkProfane(chirp string) string {
	pWords := []string{"kerfuffle", "sharbert", "fornax"}

	chirpWords := strings.Split(chirp, " ")

	for i, chirpWord := range chirpWords {

		for _, pWord := range pWords {
			if pWord == strings.ToLower(chirpWord) {
				chirpWords[i] = "****"
			}
		}
	}

	return strings.Join(chirpWords, " ")

}
