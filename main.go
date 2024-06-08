package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("assets/keyboard.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := make([]byte, 1024) 
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }

    htmlContent := fmt.Sprintf("<pre>%s</pre>", string(data[:count]))

    htmlFile, err := os.Create("ascii_art.html")
    if err != nil {
        log.Fatal(err)
    }
    defer htmlFile.Close()

    _, err = htmlFile.WriteString(htmlContent)
    if err != nil {
        log.Fatal(err)
    }
}
