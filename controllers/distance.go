package controllers

import (
	"DistanceService/helpers"
	"DistanceService/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DistanceController struct {
}

func NewDistanceController() *DistanceController {
	return &DistanceController{}
}

func (controller DistanceController) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	req := models.Request{}

	json.NewDecoder((r.Body)).Decode(&req)

	response := models.Response{}
	response.Message = "Success"
	response.Status = true
	response.Data = helpers.GetDistance(req.StartingPoint, req.DestinationPoint)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}
