package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

type UserType struct {
	data *model.User
}

func (u UserType) RegisterRoute() {
	http.HandleFunc("/record", u.RecordUserResponse)
}

func (u UserType) RecordUserResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		f := r.Form
		resp, err := model.RecordResponse(&model.User{
			Name:          f.Get("name"),
			Email:         f.Get("email"),
			Reg:           f.Get("reg"),
			ApplicantType: f.Get("applicantType"),
		})
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(resp)
	}
}
