package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/api/restful/middleware"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/domain"
)

func RegisterRoutes(app *bootstrap.Application, bookUsecase domain.BookUsecase, loginUsecase domain.LoginUsecase, signupUsecase domain.SignUpUsecase) {
	// Register Global Middleware
	cors := middleware.CORSMiddleware()
	app.Engine.Use(cors)

	// Register Book Routes
	bookController := controller.NewBookController(bookUsecase)
	RegisterBookRoutes(app.Engine, bookController)

	// Register Login Routes
	loginController := controller.NewLoginController(loginUsecase, app.Env)
	RegisterLoginRoutes(app.Engine, loginController)

	// Register Signup Routes
	signupController := controller.NewSignupController(signupUsecase)
	RegisterSignupRoutes(app.Engine, signupController)
}
