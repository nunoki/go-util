// package loading outputs an animated loading icon to the terminal 
// output, alongside a custom message.
package loading

import (
	"fmt"
	"time"
	"strings"
)

var isLoading = true
var stateLoading = 0
var frameLength = 100
var chars = []string{`\`, `|`, `/`, `—`}
var msg = ""

// Start starts the loading animation.
// This function has to be started in a goroutine.
// It will call itself until Stop() is called.
func Start(message string) {
	msg = message
	for {
		outputChar()
		if isLoading {
			wait(1)
		} else {
			break
		}
	}
}

// Stop stops the loading animation, and clears the line.
func Stop() {
	isLoading = false
	wait(1.1)
	fmt.Print("\r" + strings.Repeat(" ", len(msg) + 1) + "\r")
}

// outputChar outputs the next loading icon frame and the message.
// It will use \r to overwrite the previous frame.
func outputChar() {
	fmt.Print("\r" + chars[stateLoading] + " " + msg)
	stateLoading = (stateLoading + 1) % len(chars)
}

// wait sleeps for the configured time inbetween frames.
// The multiplier can be used to alter the default time.
func wait(multiplier float64) {
	time.Sleep(time.Duration(int(float64(frameLength) * multiplier)) * time.Millisecond)
}

