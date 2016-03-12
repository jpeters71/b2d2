package gpio

import "fmt"

// Mode is used to determine what mode a GPIO pin should be set to.
type Mode int

const (
	// Read specifies GPIO read mode should be set
	Read Mode = iota
	// Write specifies GPIO write mode should be set
	Write
)

// SetMode sets the specified mode on the given pin.
func SetMode(iPin int16, mode Mode) {
	fmt.Printf("RPI! SetMode on pin %v, mode: %v\n", iPin, mode)
}

// SetHigh sets the specified pin to high
func SetHigh(iPin int16) {
	fmt.Printf("RPI! SetHigh on pin %v\n", iPin)
}

// SetLow sets the specified pin to low
func SetLow(iPin int16) {
	fmt.Printf("RPI! SetLow on pin %v\n", iPin)
}
