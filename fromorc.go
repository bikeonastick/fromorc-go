package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func GetData(url string) []byte {
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
	resp := GetData(url)
	e2 := json.Unmarshal(resp, &t)
	if e2 != nil {
		panic(e2)
	} else {
		return
	}
}

func getTrails() string {
	resp := GetData("https://api.morcmtb.org/v1/trails")
	return string(json.RawMessage(resp))
}

func main() {
	ret := getTrailObjs("https://api.morcmtb.org/v1/trails")
	fmt.Println(ret)

}
