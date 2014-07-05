package main

type Settings struct {
	Port     int       `json : "port"`
	Hostname string    `json : "hostname"`
	Services []Service `json : "services"`
}

type Service struct {
	Name  string `json: "name"`
	Label string `json : "label"`
	Check string `json: "check"`
	Host  string `json: "host"`
	Port  int    `json: "port"`
	Path  string `json : "path"`
}

type Status struct {
	Name   string `json: "serviceName"`
	Status string `json : "serviceStatus"`
}
