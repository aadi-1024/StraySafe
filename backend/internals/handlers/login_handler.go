package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type loginJson struct {
	Username string `json:"username"`
	Password string	`json:"password"`
	//crsf token
}

type responseJson struct {
	Authenticated bool `json:"authenticated"`
	Message string	`json:"message"`

}

func SignUpPostHandler(w http.ResponseWriter, r *http.Request) {
	var reqPay loginJson
	err := json.NewDecoder(r.Body).Decode(&reqPay)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	pass_b := []byte(reqPay.Password)
	if len(pass_b) > 72 {
		log.Println("password can't be longer than 72 bytes")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = bcrypt.CompareHashAndPassword(nil, pass_b)
	if err != nil {
		log.Println(err)
	}
}