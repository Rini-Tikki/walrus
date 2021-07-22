package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PatreonUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DownloadList [][]string

var downloadList DownloadList

func UserInfo(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Println(sb)

	var user PatreonUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Couldn't parse /user body.")
		fmt.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "OK")
}

func DownloadURLCollector(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &downloadList)
	if err != nil {
		fmt.Println("Couldn't parse /download body.")
		fmt.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "OK")
}

func JSFinished(w http.ResponseWriter, req *http.Request) {
	go DownloadJobHandler(downloadList)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "OK")
}

//go:embed client.js
var jsGadget string
func ServeGadget(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, jsGadget)
}
