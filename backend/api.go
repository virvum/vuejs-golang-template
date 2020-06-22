package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// APIError issues an API error with the given error message.
func APIError(ctx *gin.Context, code int, err string) {
	log.Log(5, Error, "error: %s", err)
	ctx.JSON(code, gin.H{"error": err})
	ctx.Abort()
}

// APIStart initializes the API.
func APIStart() error {
	router := gin.New()
	router.Use(gin.Recovery())

	// Loggin middleware.
	router.Use(func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path

		ctx.Next()

		latency := time.Now().Sub(start)
		status := ctx.Writer.Status()
		l := fmt.Sprintf("%d %s %s %s %s", ctx.Writer.Status(), FormatDuration(latency, "ms"), ctx.ClientIP(), ctx.Request.Method, path)

		// TODO see https://github.com/gin-gonic/gin/blob/master/logger.go#L234
		// TODO "[GIN] 2020/04/01 - 10:26:33 | 200 |    36.39524ms |       127.0.0.1 | POST     /api/auth"
		// TODO display sent and received bytes formatted (MiB etc)

		switch {
		case status >= 200 && status < 400:
			log.Info("%s", l)
		case status >= 400 && status < 500:
			log.Warn("%s", l)
		case status >= 500:
			log.Error("%s", l)
		default:
			log.Warn("%s", l)
		}
	})

	// CORS.
	// Doesn't work when used on routerRoot instead of router. See
	// https://stackoverflow.com/a/29439630/239117 for a custom working
	// middleware.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // TODO
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	routerRoot := router.Group(*app.config.API.BasePath)

	// Public API calls (anonymous access)
	public := routerRoot.Group("/api")
	public.POST("/auth", APIAuthPost)

	// Private API calls (requires authentication)
	private := routerRoot.Group("/api")
	private.Use(APIAuthMiddleware)
	private.GET("/auth", APIAuthGet)
	private.DELETE("/auth", APIAuthDelete)

	server := &http.Server{
		Addr:           *app.config.API.Address,
		Handler:        router,
		ReadTimeout:    app.config.API.ReadTimeout,
		WriteTimeout:   app.config.API.WriteTimeout,
		MaxHeaderBytes: app.config.API.MaxHeaderBytes,
	}

	go func() {
		log.Info("starting listener on %s", *app.config.API.Address)

		if err := server.ListenAndServe(); err != nil {
			log.Fatal("http.ListenAndServe: %v", err)
			// TODO shutdown
		}
	}()

	return nil
}
