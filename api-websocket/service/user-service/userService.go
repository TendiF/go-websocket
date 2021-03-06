package userService

import (
	"os"
	"encoding/json"
	"net/http"
	"log"
	"time"
	. "go-chat/utils"
	userModel "go-chat/models/user-model"
	"go.mongodb.org/mongo-driver/mongo"
	jwt "github.com/dgrijalva/jwt-go"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method, r.URL.Path )

	if r.Method == "POST" && r.URL.Path == "/user"{
		AddUser(w, r)
		return
	}

	if r.Method == "POST" && r.URL.Path == "/user/login"{
		Login(w, r)
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

	plainPassword := user.Password

	user, err = userModel.GetByPhone(user.Phone); 
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "phone number not found", http.StatusNotAcceptable)
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if ComparePasswords(user.Password, []byte(plainPassword)) {
		atClaims := jwt.MapClaims{}
		atClaims["authorized"] = true
		atClaims["user_id"] = user.ID
		atClaims["exp"] = time.Now().Add(time.Minute * 43200).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			 log.Println(err)
		}

		b, err := json.Marshal(map[string]interface{}{
			"token": token,
			"message" : "success",
		})

		if err != nil {
			log.Println(err)
			return
		}

		w.Write([]byte(string(b)))
	} else {
		http.Error(w, "wrong phone or password", http.StatusNotAcceptable)
	}


}
