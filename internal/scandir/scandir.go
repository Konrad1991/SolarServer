package scandir

import (
  "fmt"
  "strings"
  "os"
  "log"
  "time"
  "path/filepath"
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
// update Tree due to changes 

func Scan(current *Tree, indent string, path string)  {
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

func wsiq(s *string) string {
  *s = "\"" + *s + "\""
  return *s
}

func createFileDict(indent string, head* Tree, jsonString* strings.Builder) {
  jsonString.WriteString(indent + "\"files\"" + ":" + "[")
  for i, s := range head.Files {
    jsonString.WriteString(  wsiq(&s.FileName))
    if i < (len(head.Files) - 1) {
      jsonString.WriteString(",")
    }
  }
  jsonString.WriteString("]\n")
}

func addEndParenthesis(i int, head* Tree, jsonString* strings.Builder) {
    if (i < (len(head.Nexts) - 1)) && (len(head.Nexts) > 1) {
      jsonString.WriteString("},\n")
    } else {
      jsonString.WriteString("}\n")
    }
}

func addEndForDirect(indent string, head* Tree, jsonString* strings.Builder) {
  if len(head.Nexts) == 0 {
    jsonString.WriteString(indent + "}\n")
  } else {
    jsonString.WriteString(indent + ",\n")
  }
}

func toJson(indent string, head *Tree, jsonString* strings.Builder) {
  indent = indent + " "
  jsonString.WriteString(wsiq(&head.Name) + ":" + "{\n")

  createFileDict(indent, head, jsonString)
  addEndForDirect(indent, head, jsonString);

  for i:= 0; i < len(head.Nexts); i++ {
    toJson(indent, head.Nexts[i], jsonString)
    addEndParenthesis(i, head, jsonString);
  }
}

func TreeToJson(indent string, head *Tree,
  jsonString* strings.Builder) string {
  toJson(indent, head, jsonString)
  return "{" + jsonString.String()
}






// Tobis Code
func TreeToJsonFile(head *Tree, jsonString* strings.Builder) {
jsonString.WriteString("{\n")
PrintJT(head, jsonString)
jsonString.WriteString("}")
}

func PrintJT(Tree *Tree, jsonString* strings.Builder) {
  jsonString.WriteString(wsiqt(Tree.Name) +  ":\n")
  var ArrayString strings.Builder
  for i, s := range Tree.Files {
    if i < len(Tree.Files) - 1 {
      ArrayString.WriteString(wsiqt(s.FileName) + ",\n")
    } else {
      ArrayString.WriteString(wsiqt(s.FileName) + "\n")
    }
  } 
  // s := ArrayString.String()
  // s_c := "\"files\":" + wsisb(s) // hier aufgehört

  // for i, s := range Tree.Nexts {
  //
  // }

  // jsonString.WriteString(wsicb())
}

func wsiqt(s string) string{
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

