package main

import (
	"fmt"
	"log"
	"time"
  "os"
  "html/template"
  "strconv"
)

func calcDays(from, to time.Time) int {
  duration := to.Sub(from)
  days := int(duration.Hours()/24)
  return days
}

func render(days string) {
    templateFile := "template.html"
    tmplContent, err := os.ReadFile(templateFile)
    if err != nil {
        panic(err)
    }

    data := struct {
        D1 string
        D2 string
        D3 string
    }{
        D1: string(days[0]),
        D2: string(days[1]),
        D3: string(days[2]),
    }

    t := template.Must(template.New("htmlTemplate").Parse(string(tmplContent)))

    file, err := os.Create("pages/index.html")

    if err != nil {
        panic(err)
    }
    defer file.Close()

    err = t.Execute(file, data)
    if err != nil {
        panic(err)
    }
}


func main() {
  from, err := time.Parse("2006-01-02", "2023-09-02")
  if err != nil {
    log.Fatal(err)
  }
  days := calcDays(from, time.Now())

  fmt.Printf("Days: %d", days)

  render(strconv.Itoa(days))
}

