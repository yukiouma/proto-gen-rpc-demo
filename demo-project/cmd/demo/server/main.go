package main

import (
	v1 "demo/api/demo/v1"
	"demo/internal/demo"
	"log"
)

const (
	addr    = ":8000"
	network = "tcp"
)

func main() {
	server := v1.NewServer()
	demoSvc := demo.NewDemoService()
	server.RegisterDemoService(demoSvc)
	log.Printf("[INFO]: server start and listening on: %v", addr)
	if err := server.Run(network, addr); err != nil {
		log.Fatalf("[ERROR]: server down because: %v", err)
	}
}
