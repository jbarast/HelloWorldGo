package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", indexHandle)
	http.ListenAndServe(":"+PORT, nil)
}

func indexHandle(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset = utf-8")
	fmt.Fprint(w, "<H1>Hello World</H1>")
}


