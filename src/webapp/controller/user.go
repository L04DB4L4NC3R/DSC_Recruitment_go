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
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		f := r.Form
		err = model.RecordResponse(&model.User{
			Name:          f.Get("name"),
			Email:         f.Get("email"),
			Reg:           f.Get("reg"),
			ApplicantType: f.Get("applicantType"),
		})
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("New entry added"))
		return
	}
}

func (u UserType) ShowUserResponse(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != os.Getenv("ADMIN_PASS") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Sorry, forbidden"))
		return
	}

	if r.Method == http.MethodGet {
		query := r.URL.Query()
		reg := query.Get("reg")

		if len(reg) > 0 {
			val, err := model.ShowByReg(reg)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			json.NewEncoder(w).Encode(val)
			return
		}

		val, err := model.ShowResponse()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(val)
	}

}
