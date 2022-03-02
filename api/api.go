package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anilkusc/go-package-layout/domain"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// App method is the main struct for the application
type Api struct {
	Router       *mux.Router
	SessionStore *sessions.CookieStore
	Domain       *domain.Domain
}

func (api *Api) Init() {
	api.Router = mux.NewRouter()
	api.InitRoutes()
	api.SessionStore = sessions.NewCookieStore([]byte(os.Getenv("STORE_KEY")))
}

func (api *Api) InitRoutes() {
	api.Router.HandleFunc("/factorial/calculate", api.CalculateHandler)
}

func (api *Api) Start() {
	api.Init()

	fmt.Println("Serving on: ")
	http.ListenAndServe(":8080", api.Router)
}
