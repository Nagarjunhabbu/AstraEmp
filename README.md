# API Service

This is a simple API service built with Go.

## Installation

### 1. Install Go

Ensure you have Go installed on your system. You can download and install it from the official Go website: [https://golang.org/](https://golang.org/).

### 2. Clone the Repository

Clone this repository to your local machine:

```bash
https://github.com/Nagarjunhabbu/AstraEmp.git
```
### 3.Install Docker compose

```bash
docker-compose up
```
Server will start on local host with port 8000.
Once the service is running, you can make HTTP requests to the API. Here's an example using curl:
### 4. Example API Usage
Create:
```bash
curl --location 'localhost:8000/v1/employee' \ --header 'Content-Type: application/json' \ --data '[{
	"name":"om",
	"salary":780000,
	"designation":"FE",
	"insurance_id":2,
	"insurance_amount":1000000,
	"location":"Kolar"
},
{
    "name":"Nags",
	"salary":780000,
	"designation":"Manager",
	"insurance_id":2,
	"insurance_amount":1000000,
	"location":"Mysore"

}
 ]'
```
Response
```json
{
	"message": "Employees received successfully"
}
```




```bash
curl --location 'localhost:8000/v1/employee'
```
Response
```json
[{
	"name":"om",
	"salary":780000,
	"designation":"FE",
	"insurance_id":2,
	"insurance_amount":1000000,
	"location":"Kolar"
},
{
    "name":"Nags",
	"salary":780000,
	"designation":"Manager",
	"insurance_id":2,
	"insurance_amount":1000000,
	"location":"Mysore"

}
 ]
```


