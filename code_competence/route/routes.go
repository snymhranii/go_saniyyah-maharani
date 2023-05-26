package route

import (
	"net/http"
	"project_structure/constant"
	"project_structure/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/register/user", controller.CreateUserController)
	e.POST("/login/user", controller.LoginUserController)
	e.POST("/register/admin", controller.CreateAdminController)
	e.POST("/login/admin", controller.LoginAdminController)
	e.GET("/items", controller.GetProductsController)

	admin := e.Group("/admin", middleware.JWT([]byte(constant.SECRET_JWT)))
	admin.GET("/users", controller.GetUsersController)
	user.GET("/items/:id", controller.GetPaymentController)
	admin.POST("/items", controller.CreateProductController)
	admin.PUT("/items/:id", controller.UpdateProductController)
	admin.DELETE("/items/:id", controller.DeleteProductController)
	admin.GET("/items", controller.GetOrdersController)
	admin.GET("/items/:category_id", controller.GetOrdersController)
	admin.GET("/items?keyword=item_name", controller.GetOrdersController)

}
