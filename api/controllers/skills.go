package controllers

import (
	"net/http"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/go-chi/render"
)

func GetSkills(w http.ResponseWriter, r *http.Request) {
	skills := models.GetAllSkills()
	response := models.NewResponse("Skills retrieved successfully", skills)
	render.JSON(w, r, response)
}
