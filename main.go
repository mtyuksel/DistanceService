package main

import (
	"DistanceService/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	distanceController := controllers.NewDistanceController()
	r.POST("/distance/calculate", distanceController.Calculate)

	http.ListenAndServe("localhost:9000", r)
}
