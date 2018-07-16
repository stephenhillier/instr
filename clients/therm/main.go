package main

import (
	"fmt"
	"log"
	"time"

	"github.com/stephenhillier/instr/backend/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// func client() {

// }

func main() {
	devID := "TH18-1"      // device ID
	var resistance float32 // the measured resistance of the thermistor
	// var diff float64

	// gRPC connection parameters
	var conn *grpc.ClientConn
	var err error
	host := "stevepc"
	port := 7777

	// set up the Raspberry Pi/Analog to digital converter (ADS1015)
	// board := raspi.NewAdaptor()
	// ads1015 := i2c.NewADS1015Driver(board)
	// ads1015.DefaultGain, _ = ads1015.BestGainForVoltage(5.0)

	// robot := gobot.NewRobot("thermBot",
	// 	[]gobot.Connection{board},
	// 	[]gobot.Device{ads1015, sensor}
	// )

	// Set up a gRPC connection
	// Wait until a connection is available
	log.Printf("Trying to connect to %s...", host)
connection:
	for {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%v", host, port), grpc.WithInsecure())
		if err == nil {
			log.Printf("Connected on port %v.", port)
			break
		}
		// Connection failed - wait for connection to become available.
		log.Printf("Unable to connect: %s, trying again...", err)
		time.Sleep(10 * time.Second)
	}

	defer conn.Close()

	c := api.NewResistanceClient(conn)

	var i float32
	for i = 1.0; i < 20000.0; i = i + 1.0 {
		response, err := c.ReadResistance(context.Background(), &api.ResistanceReading{Resistance: resistance + i, Device: devID})
		if err != nil {
			log.Printf("Error making ReadResistance request: %s", err)
			conn.Close()
			goto connection
		}

		log.Printf("Response from server: %s", response.Status)
	}

	// robot.Start()
}
