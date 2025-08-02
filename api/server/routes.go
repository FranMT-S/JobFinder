package server

import (
	"github.com/FranMT-S/JobFinder/controllers"
	"github.com/go-chi/chi/v5"
)

func addScraperRoutes(router chi.Router) {
	router.Post("/scraper", controllers.Scrap)
	router.Post("/scraper/remoteok", controllers.ScrapRemoteOk)
	router.Post("/scraper/workremotely", controllers.ScrapWorkRemotely)
	router.Get("/scraper/host", controllers.GetHost)
}

func addCategoriesRoutes(router chi.Router) {
	router.Get("/categories", controllers.GetCategories)
}

func addSkillsRoutes(router chi.Router) {
	router.Get("/skills", controllers.GetSkills)
}
