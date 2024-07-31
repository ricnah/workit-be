package user

import (
	"github.com/DeniesKresna/sined/config"
	"github.com/DeniesKresna/sined/service/middlewares"
	"github.com/DeniesKresna/sined/service/modules/user/handler"
	"github.com/DeniesKresna/sined/service/modules/user/usecase"
	"github.com/DeniesKresna/sined/types/constants"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase usecase.IUsecase, cfg *config.Config) {
	userHandler := handler.UserCreateHandler(userCase)

	moduleRoute := v1.Group("/user")

	moduleRoute.POST("/login", userHandler.AuthLogin)
	moduleRoute.POST("/create", userHandler.UserCreate)

	adminRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN))
	{
		adminRoute.PUT("/edit", userHandler.UserUpdate)
		adminRoute.GET("/detail/:id", userHandler.UserGetByID)
		adminRoute.POST("/search", userHandler.UserSearch)
	}

	authRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN, constants.ROLES_USER))
	{
		authRoute.POST("/get-by-email", userHandler.UserGetByEmail)
		authRoute.POST("/list-user", userHandler.UserGetAllUser)
		moduleRoute.GET("/session", userHandler.AuthSession)
	}
}

func roleCheck(userCase usecase.IUsecase, roles ...constants.Roles) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}
