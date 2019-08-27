package main

import "github.com/Azimkhan/go-protobuf-calendar/internal/configuration"

func main() {
	logger := configuration.CreateLogger("configs/logger.json")
	logger.Info("Application started.")
	// run service code goes here...
}
