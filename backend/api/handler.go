package api

import (
	"log"

	// postgres driver
	_ "github.com/lib/pq"
	"github.com/stephenhillier/instr/backend/database"
	"golang.org/x/net/context"
)

// Server is the gRPC server
type Server struct {
}

// ReadResistance reads in a resistance (ohms) from an instrument and stores it
func (s *Server) ReadResistance(ctx context.Context, in *ResistanceReading) (*ResistanceResponse, error) {

	err := database.StoreReading(in.Resistance, in.Device)
	if err != nil {
		log.Println("Database write failed")
	}

	log.Printf("[Thermistor] [%s] %v", in.Device, in.Resistance)
	return &ResistanceResponse{Status: 1}, nil
}
