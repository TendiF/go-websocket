package chatService

import (
	"strconv"
	"encoding/json"
	"net/http"
	"log"
	. "go-chat/utils"
	chatModel "go-chat/models/chat-model"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method)

	if r.Method == "GET" {
		getChat(w, r)
		return
	}

	if r.Method == "POST" {
		AddChat(w, r)
		return
	}
}

func getChat(w http.ResponseWriter, r *http.Request){
	// Example for Normal Find query
	var page int64 = 1
	var limit int64 = 10
	if keys, ok := r.URL.Query()["page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		page = i
	}

	if keys, ok := r.URL.Query()["per_page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		limit = i
	}

	data := chatModel.Get(page, limit)
	
	w.Write([]byte(data))
}

func AddChat(w http.ResponseWriter, r *http.Request){
	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
	}

	chatModel.Add(chat)

	b, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}
