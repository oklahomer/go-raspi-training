package main

import (
	"log"
	"os"
	"os/signal"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
	"syscall"
	"time"
)

// Use GPIO 3 (Physical 5) for input
var PIN = rpi.P1_5

func main() {
	// Initialize the library and load all available drivers
	_, err := host.Init()
	if err != nil {
		log.Fatalf("Failed to initialize library: %v", err)
	}

	// Setup a pin for input
	// With the simpler circuit described in the below diagram, the input level may be unstable when the switch is not pushed.
	// https://raw.githubusercontent.com/oklahomer/go-raspi-training/master/example/002-read-tactile-switch/img/diagram_internal_pulldown.png
	// Use internal pull-down register to stabilize: https://github.com/oklahomer/go-raspi-training/tree/master/example/002-read-tactile-switch#simpler-circuit-with-built-in-pull-down-resister
	//
	//   err = PIN.In(gpio.PullDown, gpio.NoEdge)
	//
	// TheGPIO pin is explicitly wired to the ground so the internal pull-down register is not required with the below diagram.
	// https://github.com/oklahomer/go-raspi-training/tree/master/example/002-read-tactile-switch#circuit-with-explicit-pull-down-mechanism
	err = PIN.In(gpio.Float, gpio.NoEdge)
	if err != nil {
		log.Fatalf("Failed to setup GPIO pin for input: %v", err)
	}

	// Receive signal to quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Keep observing until signal is sent
	log.Println("Start")
	var level gpio.Level
LOOP:
	for {
		select {
		case <-sig:
			break LOOP

		default:
			// Equivalent to "cat /sys/class/gpio/gpio3/value"
			current := PIN.Read()
			if level != current {
				log.Printf("GPIO pin level changed from %t to %t", level, current)
				level = current

				// Give a few moments to avoid bouncing.
				time.Sleep(300*time.Millisecond)
			}
		}
	}
	log.Println("Stopped")
}
