package main

import (
		"fmt"
		"net/http"
)

func main() {
	fmt.Println("Hello world");
	http.HandleFunc("/control", handler) //Will respond only to whatever the pattern is AFTER THIS DEVICE'S IP
	http.ListenAndServe(":8000", nil)

}

func handler(w http.ResponseWriter, r *http.Request){

	r.Form.Get("asdfa")	//Will interpret ?[PARAM] as input parameters and pull them from it
	fmt.Fprintf(w, "Test message")
}
