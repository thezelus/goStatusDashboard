#goStatusDashBoard

Golang implementation of status dashboard API to monitor web end points (inspired by https://github.com/statusdashboard/statusdashboard)


##Installation
- Clone the repository 
`$ git clone --bare https://github.com/thezelus/goStatusDashboard.git`

##Getting Started

- To start the server, go into the directory and type
`$ go run server.go structs.go util.go` 

- Go to [http://localhost:4000/status](http://localhost:4000/status)

- Sample response
`{"serviceStatus":[{"Name":"Couchdb server @ local","Status":"DOWN"},{"Name":"Google","Status":"UP"},{"Name":"Facebook","Status":"UP"}]}`

##Configuring services

- Use settings.json file to configure port numbers and services to monitor

````javascript
{
    "port": 4000,
    "hostname": "127.0.0.1",
    "services": [
        {
            "name": "couchdb",
            "label": "Couchdb server @ local",
            "check": "http",
            "host": "127.0.0.1",
            "port": 1234,
            "path": "/"
        }
    ]
}
````