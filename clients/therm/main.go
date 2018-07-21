package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/stephenhillier/instr/backend/api"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"google.golang.org/grpc"
)

func main() {
	devID := "RaspberryPi-1" // device ID
	// var diff float64

	// gRPC connection parameters
	var conn *grpc.ClientConn
	var err error

	host := flag.String("host", "localhost", "hostname to connect to")
	flag.Parse()

	port := 7777

	// set up the Raspberry Pi/Analog to digital converter (ADS1015)
	board := raspi.NewAdaptor()
	ads1015 := i2c.NewADS1015Driver(board)
	ads1015.DefaultGain, _ = ads1015.BestGainForVoltage(3.3)

	// Set up a gRPC connection
	// Wait until a connection is available
	log.Printf("Trying to connect to %s...", *host)
	// connection:
	for {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%v", *host, port), grpc.WithInsecure())
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

	// work is a function that collects readings and sends them to a server to be stored
	work := func() {
		gobot.Every(500*time.Millisecond, func() {

			r, _ := ads1015.ReadWithDefaults(1)
			response, err := c.ReadResistance(context.Background(), &api.ResistanceReading{Resistance: r, Device: devID})
			if err != nil {
				log.Printf("Error sending data to server")
			}
			log.Printf("Response from server: %s", response.Status)
		})
	}

	robot := gobot.NewRobot("thermBot",
		[]gobot.Connection{board},
		[]gobot.Device{ads1015},
		work,
	)

	err = robot.Start()
	if err != nil {
		log.Println(err)
	}

}
