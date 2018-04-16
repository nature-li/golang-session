package main

import (
	"net/http"
	"fmt"
	"server/session"
	"server/session/cookie"
	//"server/session/memory"
)

// hello world, the web server
var manager session.Manager

func init() {
	key := "1234567890"
	key += "1234567890"
	key += "1234567890"
	key += "12"

	var err error
	//manager, err = memory.NewManager("session_id", 60)
	manager, err = cookie.NewManager(key, "session_id", "last_access_time", 60)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

	name := s.Get("name")
	if name == nil {
		s.Set("name", "lyg")
	}
	fmt.Println("name:", name)
	fmt.Fprintln(w, "login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	manager.SessionDestroy(w, r)
	fmt.Fprintln(w, "logout")
}

func delete(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

	s.Del("name")
	fmt.Fprintln(w, "deleted")
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/delete", delete)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
	    fmt.Println(err.Error())
	}
}
