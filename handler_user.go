package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type newUser struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	newU := newUser{}
	err := decoder.Decode((&newU))
	if err != nil {
		log.Print("Error decoding json: ", err)
		return
	}

	nu, err := cfg.db.CreateUser(context.Background(), newU.Email)
	if err != nil {
		log.Print("Error creating user: ", err)
		return
	}

	user := User{
		ID:        nu.ID,
		CreatedAt: nu.CreatedAt,
		UpdatedAt: nu.UpdatedAt,
		Email:     nu.Email,
	}

	w.WriteHeader(201)
	dat, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
	}
	w.Write(dat)
}
