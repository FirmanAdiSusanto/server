package routers

import (
	// "be22/clean-arch/features/product/data"
	// "be22/clean-arch/features/product/service"
	// "be22/clean-arch/features/product/handler"
	"taskApi/app/middlewares"
	_userData "taskApi/features/user/data"
	_userHandler "taskApi/features/user/handler"
	_userService "taskApi/features/user/service"
	"taskApi/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	// factory
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	// productData := data.New(db)
	// productService := service.New(productData)
	// productHandlerAPI := handler.New(productService)

	// e.GET("/users", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "hello world",
	// 	})
	// })

	e.GET("/users", userHandlerAPI.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", userHandlerAPI.Register)
	e.DELETE("/users/:id", userHandlerAPI.Delete)
	e.GET("/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.POST("/login", userHandlerAPI.Login)
	// e.GET("/product", productHandlerAPI.GetAll)
}
