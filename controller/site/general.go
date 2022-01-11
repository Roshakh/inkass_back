package site

import (
	"html/template"
	"inkass/inkassback/database"
	"inkass/inkassback/env"
	"net/http"
)

func HomeScreen(w http.ResponseWriter, r *http.Request) {
	home := database.GetHomeScreen()
	env.Templates.ExecuteTemplate(w, "index.html", home)
}
func AboutUs(w http.ResponseWriter, r *http.Request) {
	aboutus := database.GetAboutUs()
	env.Templates.ExecuteTemplate(w, "aboutus.html", aboutus)
}

func add(x, y int) int {
	return x + y
}

func ContactScreen(w http.ResponseWriter, r *http.Request) {
	contactScreen := database.GetContactScreen()
	f := template.FuncMap{"add": add}
	env.Templates.Funcs(f)
	env.Templates.ExecuteTemplate(w, "contacts.html", contactScreen)
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
	tarifTable := database.GetTarifTable()
	f := template.FuncMap{"add": add}
	env.Templates.Funcs(f)
	env.Templates.ExecuteTemplate(w, "tarif-table.html", tarifTable)
}
func TransportValuables(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "transport-valuables.html", nil)
}
func ServiceCatalog(w http.ResponseWriter, r *http.Request) {
	service := database.GetService()
	env.Templates.ExecuteTemplate(w, "service-catalog.html", service)
}
func VacanciesItem(w http.ResponseWriter, r *http.Request) {
	env.Templates.ExecuteTemplate(w, "vacancies-item.html", nil)
}
func Vacancies(w http.ResponseWriter, r *http.Request) {
	vacancies := database.GetVacancies()
	env.Templates.ExecuteTemplate(w, "vacancies.html", vacancies)
}
