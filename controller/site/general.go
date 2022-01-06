package site

import (
	"inkass/inkassback/database"
	"inkass/inkassback/env"
	"net/http"
)

func HomeScreen(w http.ResponseWriter, r *http.Request) {
	home := database.GetHomeScreen()
	env.Templates.ExecuteTemplate(w, "index.html", home)
}
func AboutUs(w http.ResponseWriter, r *http.Request) {
	reviews := database.GetReviews()
	env.Templates.ExecuteTemplate(w, "aboutus.html", reviews)
}

func ContactScreen(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "contacts.html", nil)
}

func Leaders(w http.ResponseWriter, r *http.Request) {
	leaders := database.GetManagement()
	env.Templates.ExecuteTemplate(w, "management.html", leaders)
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "calculator.html", nil)
}
func AdressDirector(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "adress-director.html", nil)
}

func Faq(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "faq.html", nil)
}
func MapSite(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "map-site.html", nil)
}
func News(w http.ResponseWriter, r *http.Request) {
	news := database.GetNews()
	env.Templates.ExecuteTemplate(w, "news.html", news)
}
func Portfolio(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "portfolio.html", nil)
}
func SimpleNews(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "simple-news.html", nil)
}
func TarifTable(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "tarif-table.html", nil)
}
func TransportValuables(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "transport-valuables.html", nil)
}
func ServiceCatalog(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "service-catalog.html", nil)
}
func VacanciesItem(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "vacancies-item.html", nil)
}
func Vacancies(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "vacancies.html", nil)
}
