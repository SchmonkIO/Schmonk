# Schmonk Game Server

Frontend can be found [here](https://www.google.de/search?q=nil).

## Requirements

- Go 1.9+

## Compile

- `go get -d`
- `go build`

## Example Config:

`server.conf`:
```
[SERVER]
IP = "123.456.789.000"
Port = 8080
TickRate = 20
Slots = 10
CORS = false
Debug = true
```
