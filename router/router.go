package router

import (
	"github.com/ivanwang123/go-blog/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(e *echo.Echo) {

	apiGroup := e.Group("/api")

	userGroup := apiGroup.Group("/user")
	{
		userGroup.GET("", handlers.ListUsers)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.POST("/:id/update", handlers.UpdateUser)
		userGroup.POST("/:id/delete", handlers.DeleteUser)

		userGroup.POST("/register", handlers.CreateUser)
		userGroup.POST("/login", handlers.Login)
		userGroup.POST("/logout", handlers.Logout)
		userGroup.GET("/me", handlers.Me, PrivateRoute())
	}

	postGroup := apiGroup.Group("/post")
	{
		postGroup.GET("", handlers.ListPosts)
		postGroup.GET("/paginate", handlers.PaginatePosts)
		postGroup.GET("/:id", handlers.GetPost)
		postGroup.POST("", handlers.CreatePost)
		postGroup.POST("/:id/update", handlers.UpdatePost)
		postGroup.POST("/:id/delete", handlers.DeletePost)
	}

	commentGroup := apiGroup.Group("/comment")
	{
		commentGroup.GET("/post/:post-id", handlers.GetCommentsByPost)
		commentGroup.GET("/:id", handlers.GetComment)
		commentGroup.POST("", handlers.CreateComment)
		commentGroup.POST("/:id/update", handlers.UpdateComment)
		commentGroup.POST("/:id/delete", handlers.DeleteComment)
	}

	likeGroup := apiGroup.Group("/like")
	{
		likeGroup.GET("/:user-id/liked/:post-id", handlers.LikedByUser)
		likeGroup.GET("/:post-id", handlers.GetLikesByPost)
		likeGroup.POST("", handlers.ToggleLikePost)
		// likeGroup.POST("", handlers.LikePost)
		// likeGroup.POST("/unlike", handlers.UnlikePost)
	}
}

func PrivateRoute() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("SECRET123"),
		TokenLookup:   "cookie:auth_token",
	})
}
