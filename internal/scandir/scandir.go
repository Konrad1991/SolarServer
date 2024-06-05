package scandir

import (
  "fmt"
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
