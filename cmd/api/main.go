package main

import (
	"SolarServer/internal/server"
	"SolarServer/internal/scandir"
	"fmt"
)

func main() {

  scandir.Test()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
