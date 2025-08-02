package controllers

import (
	"net/http"

	"github.com/FranMT-S/JobFinder/models"
	"github.com/go-chi/render"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := models.GetCategoryAllCategories()
	response := models.NewResponse("Categories retrieved successfully", categories)
	render.JSON(w, r, response)
}
