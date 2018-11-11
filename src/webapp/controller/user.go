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
	http.HandleFunc("/show", u.ShowUserResponse)
	http.HandleFunc("/show/", u.ShowUserTypeResponse)
}

func (u UserType) RecordUserResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data := model.User{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		err = model.RecordResponse(&data)
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

func (u UserType) ShowUserTypeResponse(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path[6:]

	switch param {
	case "management", "technical", "design":
		val, err := model.ShowTypeResponse(param)
		if err != nil {
			log.Println(err)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(val)
		return
	default:
		w.Write([]byte("Sorry, could not find what you're looking for"))
		return
	}
}
