package main

import (
	"fmt"
	"log"
	"math"
)

func moveToCoordinate(x, y float64) {
	initialBoardOffset := 29.5
	moveToAbstractCoordinate(x, y+initialBoardOffset)
}

func moveToAbstractCoordinate(x, y float64) {
	var a, b float64
	lineOffset := 3.25
	WidthOfPoints := 47.5

	//fmt.Printf("%f, %f \n ", x, y)

	a = math.Sqrt(math.Pow(-lineOffset/2+x, 2) + math.Pow(y, 2))
	b = math.Sqrt(math.Pow(-lineOffset/2+WidthOfPoints-x, 2) + math.Pow(y, 2))
	moveMotorsAbsolute(a, b)
}

func moveMotorsAbsolute(aInches, bInches float64) {
	StartingA := 58.5
	StartingB := 58.0

	newA := StartingA - aInches
	newB := StartingB - bInches

	moveMotorsInches(newA, newB)
}

func moveMotorsInches(aInches, bInches float64) {
	aRev := -1.0
	bRev := -1.0
	//fmt.Printf("Inches %f %f\n", aInches, bInches)
	ticksPerInch := float64(136)
	moveTicks(int(aInches*ticksPerInch*aRev), int(bInches*ticksPerInch*bRev))
}

func moveTicks(aTicks, bTicks int) {

	fmt.Printf("%d\t%d\n", aTicks, bTicks)

	var posXH, posXL = uint8(aTicks >> 8), uint8(aTicks & 0xff) //Split aTicks into upper and lower bytes
	var posYH, posYL = uint8(bTicks >> 8), uint8(bTicks & 0xff) //Split bTicks into upper and lower bytes

	var instruction []byte //Sets up byte array instruction

	instruction = []byte{byte(1),
		byte(posXH), byte(posXL),
		byte(posYH), byte(posYL)}

	sendCommand(instruction)

}

func sendCommand(commandList []byte) {

	/**
	COMMAND LIST:
	0 -> MOVE
	1 -> DISABLE
	2 -> HOME
	*/

	//log.Printf("%v\n", commandList)

	reader.Reset(s)

	n, err := s.Write(commandList)
	if err != nil {
		log.Fatal(err)
	}

	_ = n
	//fmt.Printf("Sent %v bytes\n", n)

	reader.ReadByte()

	//fmt.Println("finished")

}
