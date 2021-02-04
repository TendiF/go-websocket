package userService

import (
	"encoding/json"
	"net/http"
	"log"
	. "go-chat/utils"
	userModel "go-chat/models/user-model"

)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method)

	if r.Method == "POST" {
		AddUser(w, r)
		return
	}

}

func AddUser(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	if user.Password == "" || len(user.Password) < 6{
		http.Error(w, "invalid password", http.StatusNotAcceptable)
		return
	}

	if user.Phone == "" {
		http.Error(w, "required Phone", http.StatusNotAcceptable)
		return
	}

	if user.Name == "" {
		http.Error(w, "required Name", http.StatusNotAcceptable)
		return
	}

	user.Password = HashAndSalt([]byte(user.Password))

	user = userModel.Add(user)

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}

func Login(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	if user.Password == "" || len(user.Password) < 6{
		http.Error(w, "invalid password", http.StatusNotAcceptable)
		return
	}

	if user.Phone == "" {
		http.Error(w, "required Phone", http.StatusNotAcceptable)
		return
	}


}
