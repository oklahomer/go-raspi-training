package main

import (
	"log"
	"os"
	"os/signal"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
	"syscall"
	"time"
)

// Check the designated address as below:
//
//   > sudo i2cdetect -y 1
//        0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
//   00:          -- -- -- -- -- -- -- -- -- -- -- -- --
//   10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
//   20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
//   30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
//   40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
//   50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
//   60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
//   70: -- -- -- -- -- -- -- --
//
// For options, see "SERIAL BUS ADDRESS" section of the data sheet at
// https://www.analog.com/media/en/technical-documentation/data-sheets/ADT7410.pdf
const Address = 0x48

const (
	// TemperatureValueRegisterAddress is the register address to read the temperature
	TemperatureValueRegisterAddress = 0x00
	// ConfigurationRegisterAddress is the register address to write the preferred configuration
	ConfigurationRegisterAddress = 0x03
)

func main() {
	// Initialize the library and load all available drivers
	_, err := host.Init()
	if err != nil {
		log.Fatalf("Failed to initialize library: %v", err)
	}

	// Get the first available I²C bus
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatalf("Failed to open I²C bus: %v", err)
	}
	defer func() {
		err = bus.Close()
		if err != nil {
			log.Printf("Failed to close I²C bus: %v", err)
		}
		log.Print("Closed I²C device")
	}()

	// Setup a I²C device on the bus
	device := &i2c.Dev{Bus: bus, Addr: Address}

	// Switch to 16-bit mode more higher resolution.
	// "The number of bits in the temperature data-word can be extended
	// to 16 bits, twos complement, by setting Bit 7 to 1 in the configuration register (Register Address 0x03)."
	err = device.Tx([]byte{ConfigurationRegisterAddress, 0x80}, nil)
	if err != nil {
		log.Printf("Failed to write to configuration register: %v", err)
	}

	// Measure temperature every 3 seconds
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println("Start")
LOOP:
	for {
		select {
		case <-ticker.C:
			write := []byte{TemperatureValueRegisterAddress}
			// Read the whole 16bits
			read := make([]byte, 2)
			err = device.Tx(write, read)
			if err != nil {
				log.Printf("Failed to read temperature: %v", err)
				continue
			}

			// 16-bit mode uses the entire 16 bits to represent the current temperature.
			tmp := uint16(read[0])<<8 | uint16(read[1])

			// Convert to human readable units
			tmpC := float32(tmp) / 128
			tmpF := tmpC*1.8 + 32

			log.Printf("Current temperature: %f°C (%f°F)", tmpC, tmpF)

		case <-sig:
			break LOOP
		}
	}
	time.Sleep(1 * time.Second) // Wait till resources are closed
	log.Println("Stopped")
}
