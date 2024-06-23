package main

import (
	"SolarServer/internal/server"
	"SolarServer/internal/scandir"
	"fmt"
  "strings"
)

func main() {
  head := &scandir.Tree{
    Name:  "/home/konrad/Documents/SolarServer",
    Files: nil,
    Nexts:  nil,
    Previous: nil,
  }
  scandir.Scan(head, "", head.Name)
  fmt.Println("\n")
  //scandir.Print("", head)
  var s strings.Builder
  scandir.TreeToJson(head, &s)
  fmt.Println(s.String())
  return;
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
