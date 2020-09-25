package routes

import (
	"net/http"

	"github.com/programmer-richa/go_rest_api/controller"
	"github.com/programmer-richa/go_rest_api/models"
	"github.com/programmer-richa/go_rest_api/settings"
	"github.com/programmer-richa/go_rest_api/utils"
)

// RequestHandler links a url with its handler
func RequestHandler(url string, method string, handler func(http.ResponseWriter, *http.Request)) {
	if method == http.MethodGet || method == http.MethodPost {
		http.HandleFunc(url, handler)
	}
}

// Router sets up API urls and dummy data of the application
func Routes() {
	settings.Users = append(settings.Users, &models.User{Email: "dummy@gmail.com", Password: "123"})
	controller := controller.Controller{}
	RequestHandler("/login", http.MethodPost, controller.Login)
	RequestHandler("/sign-up", http.MethodPost, controller.SignUp)
	RequestHandler("/dashboard", http.MethodPost, utils.TokenVerifyMiddleWare(controller.Dashboard))
}
