package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/tarm/serial"
	"bufio"
	"go.bug.st/serial.v1"
	"os"
)

var s serial.Port
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

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 { //If no ports are found
		log.Fatal("No serial ports found")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port) //Outputs "found port [PORT]" per port on a new line
	}
	mode := &serial.Mode{BaudRate: 9600}
	s, err := serial.Open("/dev/tty.usbmodem1411", mode)
	if err != nil {
		log.Fatal(err)
	}

	reader = bufio.NewReader(s)

	http.HandleFunc("/control", handler) //Will respond only to whatever the pattern is AFTER THIS DEVICE'S IP
	http.ListenAndServe(":8000", nil)
}

func control(command, posX, posY int) { //Putting data type after all variable names sets them all
	//instruction := {command, posX, posY} THIS IS BAD YOU ARE WRONG
	var posXH, posXL = uint8(posX >> 8), uint8(posX & 0xff) //Split posX into upper and lower bytes
	var posYH, posYL = uint8(posY >> 8), uint8(posY & 0xff) //Split posY into upper and lower bytes

	var instruction []byte //Sets up byte array instruction
	/**
	COMMAND LIST:
	0 -> MOVE
	1 -> DISABLE
	2 -> HOME
	*/

	switch command {
	case 0:
		instruction = []byte{byte(command),
			byte(posXH), byte(posXL),
			byte(posYH), byte(posYL)}
	case 1:
		instruction = []byte{byte(1)} //disable
	case 2:
		instruction = []byte{byte(2)}
	default:
		//cry
	}

	n, err := s.Write(instruction)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	for true {
		reader.WriteTo(os.Stdout)
	}

	reader.ReadByte() //Reads confirmation byte from Arduino, if no byte returned return error
}

func handler(w http.ResponseWriter, r *http.Request) {
	control(0, 100, 100)
	//r.Form.Get("asdfa") //Will interpret ?[PARAM] as input parameters and pull them from it
	fmt.Fprint(w, "Test message")
}
