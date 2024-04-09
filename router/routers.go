package router

import (
	"net/http"
	"recipe/config"
	"recipe/controllers"
	"recipe/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	r.Use(cors.Default())

	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	users := r.Group("/users")
	{
		users.GET("/:id", controllers.UsersController{}.GetUserByUserId)
	}

	r.GET("/recipe", controllers.RecipeController{}.GetRecipeItems)

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Wolrd!")
	})

	return r

}
