package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../model"
)

type UserType struct {
	data *model.User
}

func (u UserType) RegisterRoute() {
	http.HandleFunc("/record", u.RecordUserResponse)
	http.HandleFunc("/show", u.ShowUserResponse)
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

func (u UserType) ShowUserResponse(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != os.Getenv("ADMIN_PASS") {
		log.Println("HEYY")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Sorry, forbidden"))
		return
	}

	val, err := model.ShowResponse()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Sorry, server error"))
		return
	}
	json.NewEncoder(w).Encode(val)

}
