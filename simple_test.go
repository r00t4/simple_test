package main

import (
	"strings"
	"testing"
)

var svc = Service{users: map[string]string{
	"123": "somepass1",
	"456": "somepass2",
}}

func TestLoginWrongPass(t *testing.T) {
	req := &LoginRequest{
		GrantType: "password",
		Username:  "123",
		Password:  "somewrongpass",
	}
	_, err := svc.Login(req)
	if err != nil && !strings.Contains(err.Error(), "wrong password") {
		t.Errorf("Login(%s:%s) err = %s; want <wrong password>", req.Username, req.Password, err.Error())
	} else {
		t.Errorf("Login(%s:%s) err = nil; want <wrong password>", req.Username, req.Password)
	}
}

func TestLoginNotFound(t *testing.T) {
	req := &LoginRequest{
		GrantType: "password",
		Username:  "123",
		Password:  "somewrongpass",
	}
	_, err := svc.Login(req)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		t.Errorf("Login(%s:%s) err = %s; want <not found>", req.Username, req.Password, err.Error())
	} else {
		t.Errorf("Login(%s:%s) err = nil; want <wrong password>", req.Username, req.Password)
	}
}

func TestLoginToken(t *testing.T) {
	req := &LoginRequest{
		GrantType: "password",
		Username:  "123",
		Password:  "somewrongpass",
	}
	_, err := svc.Login(req)
	if err != nil {
		t.Errorf("Login(%s:%s) err = %s; want <nil>", err.Error())
	}
}