package main

import (
    "fmt"
    "log"
    "os"
)


func readAssest(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
      return "", nil
    }
    defer file.Close()
    data := make([]byte, 1024) 
    count, err := file.Read(data)
    if err != nil {
      return "", nil
    }
    return string(data[:count]), nil
}

func writeHtml(path string, content string) error {
    htmlContent := fmt.Sprintf("<pre>%s</pre>", content)

    htmlFile, err := os.Create(path)
    if err != nil {
      return err
    }
    defer htmlFile.Close()

    _, err = htmlFile.WriteString(htmlContent)
    if err != nil {
      return err
    }
    return nil
}

func main() {
    keyboard, err := readAssest("assets/keyboard.txt")
    if err != nil {
      log.Printf("%s", err)
      os.Exit(1)
    }
    writeHtml("keyboard.html", keyboard) 
}
