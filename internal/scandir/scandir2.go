package scandir2

import (
  "fmt"
  "os"
  "log"
  "time"
  "path/filepath"
  "strings"
)

func TreeToJson(head *Tree, jsonString* strings.Builder) {
jsonString.WriteString("{\n")
PrintJT(head, jsonString)
jsonString.WriteString("}")
}

func PrintJT(Tree *Tree, jsonString* strings.Builder) {
  jsonString.WriteString(wsiq(Tree.Name) +  ":\n")
  var ArrayString strings.Builder
  for i, s := range Tree.Files {
    if i < len(Tree.Files) - 1 {
      ArrayString.WriteString(wsiq(s.FileName) + ",\n")
    } else {
      ArrayString.WriteString(wsiq(s.FileName) + "\n")
    }
  } 
  s := ArrayString.String()
  s_c := "\"files\":" + wsisb(s) // hier aufgehört

  for i, s := range Tree.Nexts {

  }

  jsonString.WriteString(wsicb())
}

func wsiq(s string) string{
  s = "\"" + s+ "\""
  return s
}

func wsicb(s string) string{
  s = "{" + s+ "}"
  return s
}

func wsisb(s string) string{
  s = "[" + s+ "]"
  return s
}