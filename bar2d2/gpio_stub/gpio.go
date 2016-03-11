package gpio

import "fmt"

type Mode int

const (
	Read Mode = iota
	Write
)

func SetMode(iPin int16, mode Mode) {
	fmt.Printf("SetMode on pin %v, mode: %v", iPin, mode)
}

func SetHigh(iPin int16) {
	fmt.Printf("SetHigh on pin %v", iPin)
}

func SetLow(iPin int16) {
	fmt.Printf("SetLow on pin %v", iPin)
}
