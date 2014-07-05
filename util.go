package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	OK = "UP"
	KO = "DOWN"
)

func GetSettings(filename string) Settings {
	content, err := ioutil.ReadFile(filename)
	check(err)

	var settings Settings
	err = json.Unmarshal(content, &settings)
	check(err)

	return settings
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StatusMatrixGenerator(services []Service) []Status {

	statusMatrix := make([]Status, len(services))

	for i := range services {

		statusMatrix[i].Name = services[i].Label
		url := services[i].Check + "://" + services[i].Host + ":" + strconv.Itoa(services[i].Port)

		resp, err := http.Get(url)

		if err != nil || resp.StatusCode != http.StatusOK {
			statusMatrix[i].Status = KO
			continue
		}

		if resp.StatusCode == http.StatusOK {
			statusMatrix[i].Status = OK
		}
		defer resp.Body.Close()

	}

	return statusMatrix
}
