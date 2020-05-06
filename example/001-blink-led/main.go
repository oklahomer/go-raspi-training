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

// Use GPIO 23 (Physical 16) for out
var PIN = rpi.P1_16

func main() {
	// Initialize the library and load all available drivers
	_, err := host.Init()
	if err != nil {
		log.Fatalf("Failed to initialize library: %v", err)
	}

	// Blink every 3 sec
	t := time.NewTicker(3 * time.Second)
	defer t.Stop()

	// Receive signal to quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Keep blinking until signal is sent
	log.Println("Start")
	state := gpio.Low
LOOP:
	for {
		select {
		case <-t.C:
			state = !state // Flip state

			// Equivalent to "echo 1 > /sys/class/gpio/gpio23/value"
			e := PIN.Out(state)
			if e != nil {
				log.Fatalf("Failed to update GPIO state: %v", e.Error())
			}

		case <-sig:
			e := PIN.Out(gpio.Low) // Make sure to turn off
			if e != nil {
				log.Printf("Failed to update GPIO state but quitting anyway: %v", e.Error())
			}

			break LOOP
		}
	}
	log.Println("Stopped")
}
