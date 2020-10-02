package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func main() {
	handlers := HttpHandlers{Service{}}
	http.HandleFunc("/login", handlers.LoginHandler)
}

type LoginRequest struct {
	GrantType string `json:"grant_type"` //password, refresh, implicit
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpHandlers struct {
	service Service
}

func (httpHandl *HttpHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &LoginRequest{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}


	resp, err := httpHandl.service.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Service struct {
	users map[string]string
}

func (svc *Service) Login(req *LoginRequest) (*Token, error) {
	user, ok := svc.users[req.Username]
	if !ok {
		return nil, errors.New("not found")
	}
	if user != req.Password {
		return nil, errors.New("wrong password")
	}

	return &Token{
		AccessToken:  "123",
		RefreshToken: "456",
	}, nil
}