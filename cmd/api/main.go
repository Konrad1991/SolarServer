package main

import (
	"SolarServer/internal/server"
	"SolarServer/internal/scandir"
	"fmt"
  "strings"
  "os"
)

func main() {
  head := &scandir.Tree{
    Name:  "/home/konrad/Documents/ast2ast",
    Files: nil,
    Nexts:  nil,
    Previous: nil,
  }
  scandir.Scan(head, "", head.Name)
   scandir.Print("", head)

  var sb strings.Builder
  res := scandir.TreeToJson("", head, &sb)
  fmt.Println(res)
  d1 := []byte(res)
  os.WriteFile("/home/konrad/Documents/testFile.json", d1, 0644)
  return;
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
