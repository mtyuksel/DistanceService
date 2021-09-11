## DistanceService
Service finds the distance between two coordinates in kilometers, meters and miles.

### Running the service

```go
go build 
go run main.go
```

#### Request

- StartingPoint   : Istanbul (Turkey)
- DestinationPoint: Ankara (Turkey)

```json
{
    "StartingPoint": {
        "Latitude": 41.015137,
        "Longitude": 28.979530
    },
    "DestinationPoint": {
        "Latitude": 39.9035557,
        "Longitude": 32.6226798
    }
}
```

#### Response

```json
{
    "Status": true,
    "Message": "Success",
    "Data": {
        "Kilometer": 332.05173021062205,
        "Mile": 206.32731565270444,
        "Meter": 332051.7302106221
    }
}
```
