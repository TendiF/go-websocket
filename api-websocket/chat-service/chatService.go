package chatService

import (
	"net/http"
)

func GetChat(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}
