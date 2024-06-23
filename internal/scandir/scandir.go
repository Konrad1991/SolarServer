package scandir

import (
  "fmt"
  "os"
  "log"
  "time"
  "path/filepath"
  "strings"
)

type File struct {
  FileName string
  Size int64
  Date time.Time
  Extension string
}

type Tree struct {
  Name string
  Files []File
  Previous *Tree
  Nexts []*Tree
}

// TODO: 
// size method for Tree
// update ree due to changes 

func Scan(current *Tree, indent string, path string)  {
  fmt.Println(indent + path)
  entries, err := os.ReadDir(path)
  if err != nil {
    log.Fatal(err)
  }

  for _, e := range entries {
    if e.IsDir() {
      newNode := &Tree{
        Name:     e.Name(),
        Previous: current,
      }
      current.Nexts = append(current.Nexts, newNode)
      Scan(newNode, indent + "\t", path + "/" + e.Name())
    } else {
      fileInfo, err := e.Info()
      if err != nil {
        log.Fatal(err)
      }
      newFile := &File{
        FileName: e.Name(),
        Size: fileInfo.Size(),
        Date: fileInfo.ModTime(),
        Extension: filepath.Ext(e.Name()),
      }
      current.Files = append(current.Files, *newFile)
      fmt.Println(indent + e.Name())
    }
  }
}

func Print(indent string, head *Tree) {
  indent = indent + " "
  fmt.Println(head.Name)
  for i, s := range head.Files {
    fmt.Println(indent + "|", indent + "|--", i,
      " | ", s.Size, " bp | ",
      s.FileName)
  }
  for i:= 0; i < len(head.Nexts); i++ {
    Print(indent, head.Nexts[i])
  }
}

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
  jsonString.WriteString(wsicb("\"files\":" + wsisb(s)))
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