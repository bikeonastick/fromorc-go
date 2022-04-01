package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://api.morcmtb.org/v1/trails"

type Trails []struct {
	TrailName   string `json:"trailName"`
	TrailID     string `json:"trailId"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedBy   string `json:"updatedBy"`
	State       string `json:"state"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	TrailStatus string `json:"trailStatus"`
	UpdatedAt   int64  `json:"updatedAt"`
	Longitude   string `json:"longitude"`
	//Description interface{} `json:"description"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Street      string `json:"street"`
}

type model struct {
	data Trails
	err  string
}

func checkServer() tea.Msg {
	//currently very happy path centric
	trails := getTrailObjs(url)

	if len(trails) == 0 {
		return errMsg{"no trails found"}
	}

	return dataMsg(trails)
}

type dataMsg Trails

type errMsg struct{ err string }

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataMsg:
		m.data = Trails(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg.err
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m model) View() string {
	if m.err != "" {
		return fmt.Sprintf("\nThere was an error: %v\n\n", m.err)
	}

	s := fmt.Sprintf("checking %s ...", url)

	for _, trail := range m.data {
		s += fmt.Sprintf(fmt.Sprintf("%s - %s - %s\n", trail.TrailName,
			trailStatusEmoji(trail.TrailStatus),
			trailStatusConfidenceEmoji(trail.UpdatedAt)))
	}

	return "\n" + s + "\n\n"
}

func trailStatusEmoji(status string) string {
	if status == "Closed" {
		return "\xF0\x9F\x91\x8E"
	} else if status == "Open" {
		return "\xF0\x9F\x91\x8D"
	} else {
		return fmt.Sprintf("¯\\_(ツ)_/¯ - %s", status)
	}

}

func trailStatusConfidenceEmoji(updatedAt int64) string {
	updatedTime := time.UnixMilli(updatedAt)
	var timeSince time.Duration = time.Now().Sub(updatedTime)
	if timeSince.Hours() <= 48 {
		return "\xe2\x9c\x85"
	} else if timeSince.Hours() <= 168 {
		return "\xf0\x9f\xa4\x9e"
	} else {
		return "\xf0\x9f\x92\xa9"
	}
}

func getData(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	var client = http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return jsonByte
}

func getTrailObjs(url string) (t Trails) {
	resp := getData(url)
	/*e2 :=*/ json.Unmarshal(resp, &t)
	//TODO: bubble possible data errors up to the view
	return t
}

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}

}
