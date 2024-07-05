package main

import (
	"fmt"
	"log"
	"time"
)

func calcDays(from, to time.Time) int {
  duration := to.Sub(from)
  days := int(duration.Hours()/24)
  return days
}

func main() {
  from, err := time.Parse("2006-01-02", "2023-09-02")
  if err != nil {
    log.Fatal(err)
  }
  days := calcDays(from, time.Now())

  fmt.Printf("Days: %d", days)
}

