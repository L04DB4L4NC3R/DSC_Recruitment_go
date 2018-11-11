package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../middleware"
)

var (
	u UserType
	m ManagementType
)

type Admin struct {
	Username string `json:username,omitempty`
	Password string `json:password,omitempty`
}

func Startup() {
	u.RegisterRoute()
	m.registerRoute()
	http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {
		a := Admin{}
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			log.Println(err.Error())
			return
		}
		if a.Username == os.Getenv("ADMIN_USERNAME") && a.Password == os.Getenv("ADMIN_PASSWORD") {
			token := middleware.GenerateJWT(os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"))
			w.Write([]byte(token))
		} else {
			w.Write([]byte("Wrong credentials"))
		}
	})
}
