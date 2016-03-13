// +build 386 amd64

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

// Init should be called before any other GPIO operations
func Init() {
	fmt.Printf("Init() called\n")
}

// SetMode sets the specified mode on the given pin.
func SetMode(iPin int16, mode Mode) {
	fmt.Printf("SetMode on pin %v, mode: %v\n", iPin, mode)
}

// SetHigh sets the specified pin to high
func SetHigh(iPin int16) {
	fmt.Printf("SetHigh on pin %v\n", iPin)
}

// SetLow sets the specified pin to low
func SetLow(iPin int16) {
	fmt.Printf("SetLow on pin %v\n", iPin)
}

// Destroy should be called when you're done with GPIO operations.  Once
// Destroy is called, Init() must be called again before more GPIO operations.
func Destroy() {
	fmt.Printf("Destroy called\n")
}
