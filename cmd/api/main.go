package main

import (
	"SolarServer/internal/server"
	"SolarServer/internal/scandir"
	"fmt"
)

func main() {
  head := &scandir.Tree{
    Name:  "/home/konrad/Documents/ETR",
    Files: nil,
    Nexts:  nil,
    Previous: nil,
  }
  scandir.Scan(head, "", "/home/konrad/Documents/ETR")
  fmt.Println("\n")
  scandir.Print("", head)
  return;
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
