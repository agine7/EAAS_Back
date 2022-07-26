package handlers

import "eaas_back/app/models"

type errorRes struct {
	Error    string `json:"error"`
	Code     int    `json:"code"`
	ErrorDis string `json:"error_discription"`
}

type basicResponse struct {
	Message string `json:"message"`
}

type loginRes struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type signupReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type submissionReq struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
