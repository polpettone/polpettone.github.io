package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"
)

func calcDays(from, to time.Time) int {
	duration := to.Sub(from)
	days := int(duration.Hours() / 24)
	return days
}

func render_days(days string) {
	templateFile := "templates/days-keyboard.html"
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

func render_kilometer(kilometers string) {
	templateFile := "templates/kilometer-keyboard.html"
	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}

	data := struct {
		D1 string
		D2 string
		D3 string
		D4 string
	}{
		D1: string(kilometers[0]),
		D2: string(kilometers[1]),
		D3: string(kilometers[2]),
		D4: string(kilometers[2]),
	}

	t := template.Must(template.New("htmlTemplate").Parse(string(tmplContent)))

	file, err := os.Create("pages/kilometer-keyboard.html")

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

	render_days(strconv.Itoa(days))
	render_kilometer("3830")
}
