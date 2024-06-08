package main

import (
    "fmt"
    "log"
    "os"
    "time"
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

func prepareKeyboard(keyboard string, value int) string  {
  return fmt.Sprintf(keyboard, value, 2, 3)
}

func daysSinceGivenDate(year int, month time.Month, day int) int {
    givenDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
    currentDate := time.Now()

    days := int(currentDate.Sub(givenDate).Hours() / 24)
    return days
}

func main() {
    keyboard, err := readAssest("assets/keyboard.txt")
    if err != nil {
      log.Printf("%s", err)
      os.Exit(1)
    }
    daysSince := daysSinceGivenDate(2023, 9, 2)
    keyboard = prepareKeyboard(keyboard, daysSince)
    writeHtml("keyboard.html", keyboard) 
}
