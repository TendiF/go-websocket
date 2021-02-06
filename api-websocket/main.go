// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	chatService "go-chat/service/chat-service"
	userService "go-chat/service/user-service"
	utilsService "go-chat/utils"
	. "go-chat/service/websocket-service"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8081", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	utilsService.CreateConnection()
	flag.Parse()
	hub := NewHub()
	go hub.Run()
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	s := router.PathPrefix("/user").Subrouter()
		s.HandleFunc("", userService.Main)
		s.HandleFunc("/login", userService.Login)
	router.HandleFunc("/chat", chatService.Main)
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8081", router))
}
