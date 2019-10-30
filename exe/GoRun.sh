mosquitto -v -p 8080 &

go build ../cmd/api/main.go
go build ../cmd/subscriber/BY_AIRPORT/AirportA/subscriberA.go
go build ../cmd/subscriber/BY_AIRPORT/AirportB/subscriberB.go
go build ../cmd/publisher/pressure/pressure.go
go build ../cmd/publisher/temperature/temperature.go
go build ../cmd/publisher/wind/wind.go

./api.exe &
./subscriberA.exe &
./subscriberB.exe &
./pressure.exe &
./temperature.exe &
./wind.exe