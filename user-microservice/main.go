package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)
type usersMap struct {
	sync.RWMutex
	users map[unit]User
}
var users usersMap // users by ID - NOT THREADSAFE
// var mapChan chan func()


func main() {

	users = usersMap{
		users: make[map[unit]User]
	}
	// users = make(map[uint]User)
  // mapchan = make(chan func())
	//
  // go func() {
	// 	for f := range mapchan {
	// 		f()
	// 	}
	// }()

	http.HandleFunc("/", getHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/delete", deleteHandler)
	fmt.Println(http.ListenAndServe("0.0.0.0:8088", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idnum, err := strconv.Atoi(id)
	if err != nil || idnum < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

  // var user User
  // var ok bool
	// getuser := func() {
	// 	user, ok = users[uint(idnum)]
	// }
	//
	// mapchan <- getuser
	users.RLock()
	defer users.Unlock()
	user, ok := users.users[uint(idnum)]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	d, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(d)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.Lock()
	defer users.Unlock()
	users.users[user.ID] = user
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idnum, err := strconv.Atoi(id)
	if err != nil || idnum < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users.Lock()
	defer users.Unlock()
	if _, ok := users.users[uint(idnum)]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(users, uint(idnum))
}
