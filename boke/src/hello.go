package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		f, err := ioutil.ReadFile("../html/Hello.html") 
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s", string(f))
	})

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
