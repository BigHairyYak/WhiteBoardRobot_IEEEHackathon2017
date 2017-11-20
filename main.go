package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"strconv"
	"time"
)

var s *serial.Port

var reader *bufio.Reader

func main() {
	//fmt.Println("Hello world");
	//c := &serial.Config{Name: "COM2", Baud: 9600}
	//s, err := serial.OpenPort(c)//open serial port
	//
	//if err != nil {
	//	log.Fatal(err)			//log error
	//}
	//
	//n, err := s.Write([]byte("test"))
	//if err != nil{
	//	log.Fatal(err)
	//}
	//
	//buf := make([]byte, 128)	//Make a 128-byte buffer
	//n, err = s.Read(buf)		//Try to read something
	//if err != nil {			//If any error is returned
	//	log.Fatal(err)			//log error
	//}
	//log.Printf("%q", buf[:n])

	c := &serial.Config{Name: "/dev/tty.usbmodem1411", Baud: 9600}
	var err error

	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)

	reader = bufio.NewReader(s)

	//control(1, -1000, -1000)
	//moveMotorsInches(0, 10)
	//moveToCoordinate(-48, -20)

	http.HandleFunc("/draw", drawRequest)
	http.HandleFunc("/command_list", commandListHandler)
	http.HandleFunc("/control", htmlHandler) //Will respond only to whatever the pattern is AFTER THIS DEVICE'S IP
	http.HandleFunc("/", positionResponse)
	http.ListenAndServe(":8000", nil)
}

func drawRequest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index2.html")
}

func positionResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {

		r.ParseForm()
		xposstr := r.PostFormValue("x-pos")
		yposstr := r.PostFormValue("y-pos")
		xpos := float64(0)
		ypos := float64(0)
		var err error
		if xposstr != "" {
			xpos, err = strconv.ParseFloat(xposstr, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		if yposstr != "" {
			ypos, err = strconv.ParseFloat(yposstr, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("x position = ", xpos)
		fmt.Println("y position = ", ypos)
		moveToCoordinate(float64(xpos), float64(ypos))
		htmlHandler(w, r)
	} else {
		htmlHandler(w, r)
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
	//control(0, 100, 100)
	//r.Form.Get("asdfa") //Will interpret ?[PARAM] as input parameters and pull them from it
	///	fmt.Fprint(w, "Test message")
}
