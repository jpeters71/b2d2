// +build arm

package gpio

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

// Mode is used to determine what mode a GPIO pin should be set to.
type Mode int

const (
	// Read specifies GPIO read mode should be set
	Read Mode = iota
	// Write specifies GPIO write mode should be set
	Write
)

// Init should be called before any other GPIO operations
func Init() {
	if err := rpio.Open(); err != nil {
		fmt.Printf("ERROR openning rpio:  %v\n", err)
	}
}

// SetMode sets the specified mode on the given pin.
func SetMode(iPin int16, mode Mode) {
	fmt.Printf("RPI! SetMode on pin %v, mode: %v\n", iPin, mode)

	pin := rpio.Pin(iPin)
	switch mode {
	case Read:
		pin.Input()
	case Write:
		pin.Output()
	}
}

// SetHigh sets the specified pin to high
func SetHigh(iPin int16) {
	pin := rpio.Pin(iPin)
	pin.High()
}

// SetLow sets the specified pin to low
func SetLow(iPin int16) {
	pin := rpio.Pin(iPin)
	pin.Low()
}

// Destroy should be called when you're done with GPIO operations.  Once
// Destroy is called, Init() must be called again before more GPIO operations.
func Destroy() {
	rpio.Close()
}
