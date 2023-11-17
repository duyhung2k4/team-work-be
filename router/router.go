package router

import (
	"log"
	"net/http"
	"team-work-be/config"
	"team-work-be/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Router() http.Handler {
	app := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	app.Use(cors.Handler)
	app.Use(middleware.Logger)

	accessControler := controller.NewAccessController()
	basicQueryController := controller.NewBasicQueryController()
	advanceFilterController := controller.NewAdvanceFilterController()
	registerController := controller.NewRegisterController()

	app.Route("/api/v1", func(r chi.Router) {
		r.Get("/test-api", func(w http.ResponseWriter, r *http.Request) {
			data := controller.Response{
				Data:    "hehe",
				Message: "OK",
				Success: true,
				Error:   "",
			}

			render.JSON(w, r, data)
		})

		r.Route("/access", func(accessR chi.Router) {

			accessR.Post("/login", accessControler.Login)

			accessR.Route("/register", func(registerR chi.Router) {
				registerR.Post("/send-info-register", registerController.SendInfoRegister)
				registerR.Post("/send-code", registerController.SendCodeRegister)
			})

		})

		r.Route("/", func(protected chi.Router) {
			// protected.Use(jwtauth.Verifier(config.GetJWT()))
			// protected.Use(jwtauth.Authenticator)

			protected.Route("/access-protected", func(aProtected chi.Router) {
				aProtected.Post("/login-token", accessControler.LoginWithToken)
			})

			protected.Route("/basic-query", func(basicQueryR chi.Router) {
				basicQueryR.Post("/query", basicQueryController.BasicQuery)
			})

			protected.Route("/advance-filter", func(advanceFilterR chi.Router) {
				advanceFilterR.Post("/filter", advanceFilterController.AdvanceFilter)
			})

		})
	})

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.GetAppPort()+"/swagger/doc.json"),
	))

	log.Println("http://localhost:" + config.GetAppPort())
	log.Println("http://localhost:" + config.GetAppPort() + "/swagger/index.html")

	return app
}
