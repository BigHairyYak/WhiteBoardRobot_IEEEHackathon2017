package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func commandListHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Could not recieve list - %s", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	coords := make([]WhiteboardCoordiante, 0)
	err = json.Unmarshal(bytes, &coords)
	if err != nil {
		fmt.Printf("Could not parse JSON - %s", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(coords)

	for _, v := range coords {
		moveToCoordinate(v.X, v.Y)
	}
	fmt.Fprint(w, "Finished List")

}

type WhiteboardCoordiante struct {
	X, Y float64
}

func (c *WhiteboardCoordiante) String() string {
	return fmt.Sprintf("(%f, %f)", c.X, c.Y)
}
