package router

import (
	"html/template"
	controller "inkass/inkassback/controller"
	sc "inkass/inkassback/controller/site"
	"inkass/inkassback/env"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() {
	env.Router = mux.NewRouter()

	env.Templates = template.Must(template.ParseGlob("./templates/*.html"))
	env.Router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	env.Router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))
	env.Router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./templates/img"))))
	env.Router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./templates/fonts"))))
	// env.Router.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	env.Router.HandleFunc("/", sc.HomeScreen)
	env.Router.HandleFunc("/index.html", sc.HomeScreen)
	env.Router.HandleFunc("/aboutus.html", sc.AboutUs)
	env.Router.HandleFunc("/contacts.html", sc.ContactScreen)
	env.Router.HandleFunc("/management.html", sc.Leaders)
	env.Router.HandleFunc("/adress-director.html", sc.AdressDirector)
	env.Router.HandleFunc("/calculator.html", sc.Calculator)
	env.Router.HandleFunc("/faq.html", sc.Faq)
	env.Router.HandleFunc("/map-site.html", sc.MapSite)
	env.Router.HandleFunc("/news.html", sc.News)
	env.Router.HandleFunc("/portfolio.html", sc.Portfolio)
	env.Router.HandleFunc("/service-catalog.html", sc.ServiceCatalog)
	env.Router.HandleFunc("/simple-news.html", sc.SimpleNews)
	env.Router.HandleFunc("/tarif-table.html", sc.TarifTable)
	env.Router.HandleFunc("/transport-valuables.html", sc.TransportValuables)
	env.Router.HandleFunc("/vacancies-item.html", sc.VacanciesItem)
	env.Router.HandleFunc("/vacancies.html", sc.Vacancies)
	// env.Router.Handle("/templates", http.FileServer(http.Dir("./templates")))

	apiRouter := env.Router.PathPrefix("/api/").Subrouter()

	//*********HomeScreen

	apiRouter.HandleFunc("/home", controller.GetHomeScreen).Methods("GET")
	apiRouter.HandleFunc("/home/{id}", controller.DeleteHomeScreen).Methods("DELETE")
	apiRouter.HandleFunc("/home/{id}", controller.EditHomeScreen).Methods("PUT")
	apiRouter.HandleFunc("/home", controller.CreateHomeScreen).Methods("POST")

	//*********ContactScreen

	apiRouter.HandleFunc("/contactScreen", controller.GetContactScreen).Methods("GET")
	apiRouter.HandleFunc("/contactScreen/{id}", controller.DeleteContactScreen).Methods("DELETE")
	apiRouter.HandleFunc("/contactScreen/{id}", controller.EditContactScreen).Methods("PUT")
	apiRouter.HandleFunc("/contactScreen", controller.CreateContactScreen).Methods("POST")

	//*********Reviews

	apiRouter.HandleFunc("/reviews", controller.GetReviews).Methods("GET")
	apiRouter.HandleFunc("/reviews/{id}", controller.DeleteReviews).Methods("DELETE")
	apiRouter.HandleFunc("/reviews/{id}", controller.EditReviews).Methods("PUT")
	apiRouter.HandleFunc("/reviews", controller.CreateReviews).Methods("POST")

	//*********Management

	apiRouter.HandleFunc("/management", controller.GetManagement).Methods("GET")
	apiRouter.HandleFunc("/management/{id}", controller.DeleteManagement).Methods("DELETE")
	apiRouter.HandleFunc("/management/{id}", controller.EditManagement).Methods("PUT")
	apiRouter.HandleFunc("/management", controller.CreateManagement).Methods("POST")

	//*********News

	apiRouter.HandleFunc("/news", controller.GetNews).Methods("GET")
	apiRouter.HandleFunc("/news/{id}", controller.DeleteNews).Methods("DELETE")
	apiRouter.HandleFunc("/news/{id}", controller.EditNews).Methods("PUT")
	apiRouter.HandleFunc("/news", controller.CreateNews).Methods("POST")

	//*********Vacancies

	apiRouter.HandleFunc("/vacancies", controller.GetVacancies).Methods("GET")
	apiRouter.HandleFunc("/vacancies/{id}", controller.DeleteVacancies).Methods("DELETE")
	apiRouter.HandleFunc("/vacancies/{id}", controller.EditVacancies).Methods("PUT")
	apiRouter.HandleFunc("/vacancies", controller.CreateVacancies).Methods("POST")

	//*********AboutUs

	apiRouter.HandleFunc("/aboutUs", controller.GetAboutUs).Methods("GET")
	apiRouter.HandleFunc("/aboutUs/{id}", controller.DeleteAboutUs).Methods("DELETE")
	apiRouter.HandleFunc("/aboutUs/{id}", controller.EditAboutUs).Methods("PUT")
	apiRouter.HandleFunc("/aboutUs", controller.CreateAboutUs).Methods("POST")

	//*******UPLOAD/DOWNLOAD
	apiRouter.HandleFunc("/file-upload", controller.MultipleFileUpload).Methods("POST")
	apiRouter.HandleFunc("/file-download/{name}", controller.GetUploadedFiles).Methods("GET")
}
