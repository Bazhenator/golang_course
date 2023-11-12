package main

import (
  "bytes"
  "embed"
  "fmt"
  "image"
  "image/png"
  "log"
  "os"
)

//go:embed testdata/hello.txt
var hello string

//go:embed testdata/image.png
var embImage []byte

//go:embed testdata/txt
var dir embed.FS

func main() {
  //example with string
  log.Println(string(hello))

  //example with []byte
  img, _, err := image.Decode(bytes.NewReader(embImage))
  if err != nil {
    log.Fatal(err)
  }

  outFilePath := "testdata/embedded_image.png"
  outFile, err := os.Create(outFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer outFile.Close()

  err = png.Encode(outFile, img)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Image has been saved in %s\n", outFilePath)

  //example with embed.FS
  entries, _ := dir.ReadDir("testdata/txt")
  for _, entry := range entries {
    log.Println(entry.Name() + ": ")
    str, _ := dir.ReadFile("testdata/txt/" + entry.Name())
    log.Print(string(str))
  }
}
