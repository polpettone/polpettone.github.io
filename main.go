package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

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

func fetchAPIData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func fetchRunDataFromSmashrun(smashrunURL string, smashrunToken string) (*RunningStats, error) {
	fetchUrl := fmt.Sprintf("%s?access_token=%s", smashrunURL, smashrunToken)
	rawData, err := fetchAPIData(fetchUrl)

	if err != nil {
		return nil, err
	}

	var stats RunningStats
	err = json.Unmarshal([]byte(rawData), &stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

type RunningStats struct {
	RunCount                       int      `json:"runCount"`
	TotalDistance                  float64  `json:"totalDistance"`
	AverageSpeed                   float64  `json:"averageSpeed"`
	AveragePace                    string   `json:"averagePace"`
	LongestRun                     float64  `json:"longestRun"`
	LongestRunWhen                 string   `json:"longestRunWhen"`
	LongestBreakBetweenRuns        int      `json:"longestBreakBetweenRuns"`
	LongestBreakBetweenRunsDate    string   `json:"longestBreakBetweenRunsDate"`
	LongestStreak                  int      `json:"longestStreak"`
	LongestStreakDate              string   `json:"longestStreakDate"`
	AverageDistancePerDay          *float64 `json:"averageDistancePerDay"`
	AverageRunLength               float64  `json:"averageRunLength"`
	AverageDaysRunPerWeek          float64  `json:"averageDaysRunPerWeek"`
	MostOftenRunOnDay              string   `json:"mostOftenRunOnDay"`
	MostOftenRunOnCount            int      `json:"mostOftenRunOnCount"`
	MostOftenRunOnAverageDistance  float64  `json:"mostOftenRunOnAverageDistance"`
	LeastOftenRunOnDay             string   `json:"leastOftenRunOnDay"`
	LeastOftenRunOnCount           int      `json:"leastOftenRunOnCount"`
	LeastOftenRunOnAverageDistance float64  `json:"leastOftenRunOnAverageDistance"`
	DaysRunAM                      int      `json:"daysRunAM"`
	DaysRunPM                      int      `json:"daysRunPM"`
	DaysRunBoth                    int      `json:"daysRunBoth"`
}

func main() {

	apiToken := os.Getenv("SMASHRUN_API_TOKEN")

	runningStats, err := fetchRunDataFromSmashrun("https://api.smashrun.com/v1/my/stats", apiToken)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	days := runningStats.LongestStreak
	totalDistance := int(runningStats.TotalDistance)
	totalDistanceStr := strconv.Itoa(totalDistance)

	fmt.Printf("%d \n", totalDistance)
	fmt.Printf("%d \n", days)

	render_days(strconv.Itoa(days))
	render_kilometer(totalDistanceStr)
}
