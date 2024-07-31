package service

import (
	"github.com/DeniesKresna/sined/config"
	"github.com/DeniesKresna/sined/service/middlewares"

	userModule "github.com/DeniesKresna/sined/service/modules/user"
	userrepo "github.com/DeniesKresna/sined/service/modules/user/repository"
	usercase "github.com/DeniesKresna/sined/service/modules/user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRoutes(cfg *config.Config) (r *gin.Engine, err error) {
	r = gin.New()

	userRepo := userrepo.UserCreateRepository(cfg.DB)
	userCase := usercase.UserCreateUsecase(userRepo)

	r.Use(corsConfig())
	r.Use(middlewares.ActivityLogger())

	// Group API routes under /api
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		userModule.InitRoutes(v1, userCase, cfg)
	}

	// Redirect root URL to /public
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/public")
	})

	// Serve static files from the public directory with the updated prefix
	r.Static("/public", "./view/.output/public")

	// Catch-all route to serve index.html for client-side routing
	r.NoRoute(func(c *gin.Context) {
		c.File("./view/.output/public/index.html")
	})

	return
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include "Content-Type" in the list of allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}

func Start(cfg *config.Config) (err error) {
	eng, err := setRoutes(cfg)

	eng.Run(cfg.App.Host + ":" + cfg.App.Port)
	return
}
