package helpers

import (
	"DistanceService/models"
	"math"
)

func GetDistance(start, destination models.Location) models.Distance {
	const earthRadius = 6371

	dLat := degreeToRadian(destination.Latitude - start.Latitude)
	dLon := degreeToRadian(destination.Longitude - start.Longitude)

	lat1 := degreeToRadian(start.Latitude)
	lat2 := degreeToRadian(destination.Latitude)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	kilometer := earthRadius * c

	return models.Distance{Kilometer: kilometer, Meter: kilometer * 1000, Mile: kilometer * 0.621371}
}

func degreeToRadian(deg float64) float64 {
	return deg * math.Pi / 180
}
