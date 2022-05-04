package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

type userDetail struct {
	User string
	Text string
}

type chatDetail struct {
	Time int64
	userDetail
}

type chatDb []chatDetail

func main() {
	//var abc = userDetail{"gaurav", "my name"}
	//var cd = chatDetail{abc, 12345}
	cdb := chatDb{}
	var ChatMutex = sync.Mutex{}

	http.HandleFunc("/status", func(c http.ResponseWriter, req *http.Request) {
		c.Write([]byte("alive"))
	})

	http.HandleFunc("/message", func(w http.ResponseWriter, req *http.Request) {
		var usr userDetail
		ChatMutex.Lock()
		err := json.NewDecoder(req.Body).Decode(&usr)
		if err != nil {
			log.Printf("error in decoding")
		}
		unixTimestamp := chatDetail{time.Now().Unix(), usr}
		cdb = append(cdb, unixTimestamp)
		fmt.Fprint(w, "{\n 'ok': true\n}")
		ChatMutex.Unlock()
	})

	http.HandleFunc("/messages", func(w http.ResponseWriter, req *http.Request) {
		//var dbinst chatDb
		ChatMutex.Lock()
		sort.Slice(cdb, func(k, l int) bool { return cdb[l].Time < cdb[k].Time })
		msg := make([]chatDetail, 100)
		numberOfMsg := copy(msg, cdb)
		messagesDetail, err := json.MarshalIndent(msg[:numberOfMsg], "", " ")
		if err != nil {
			log.Printf("error in marshalling")
		}
		fmt.Fprintf(w, "{\"messages\":   %+v}", string(messagesDetail))
		ChatMutex.Unlock()
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		ChatMutex.Lock()
		userName := []string{}
		for _, us := range cdb {
			userName = append(userName, us.User)
		}
		userOut, err := json.MarshalIndent(userName, "", " ")
		if err != nil {
			log.Printf("error in marshalling of user %s", err)
		}

		fmt.Fprintf(w, "{ \"users\":%+v}", string(userOut))
		ChatMutex.Unlock()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Error 404 %s", req.URL)
	})

	err := http.ListenAndServe(":8081", nil)
	panic(err)

}
