package server

import (
	"github.com/gin-gonic/gin"
	"nunu-project/internal/handler"
	"nunu-project/internal/middleware"
	"nunu-project/pkg/helper/resp"
	"nunu-project/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler,
	orderHandler handler.OrderHandler,
	blogHandler handler.BlogHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// 无权限路由
	noAuthRouter := r.Group("/").Use(middleware.RequestLogMiddleware(logger))
	{
		noAuthRouter.GET("/order", orderHandler.GetOrderById)  // order
		noAuthRouter.GET("/blog", blogHandler.GetBlogById)     // blog
		noAuthRouter.GET("/blogList", blogHandler.GetBlogList) // blog
		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]string{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// 非严格权限路由
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// 严格权限路由
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		strictAuthRouter.GET("/users", userHandler.GetList)
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
