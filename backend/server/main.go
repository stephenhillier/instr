package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/stephenhillier/instr/backend/api"
	"github.com/stephenhillier/instr/backend/database"
	"google.golang.org/grpc"
)

func main() {
	port := 7777
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")

	database.InitDB(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbuser, dbpass, dbhost, dbname))
	defer database.DB.Close()

	s := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterResistanceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	} else {
		log.Printf("Server running at port %v", port)
	}
}
