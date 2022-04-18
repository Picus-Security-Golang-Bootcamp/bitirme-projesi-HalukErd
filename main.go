package main

import (
	_ "BasketProjectGolang/docs"
	"BasketProjectGolang/internal/service"
	"BasketProjectGolang/pkg/config"
	db "BasketProjectGolang/pkg/database"
	"BasketProjectGolang/pkg/graceful"
	logger "BasketProjectGolang/pkg/logging"
	mw "BasketProjectGolang/pkg/middleware"
	"BasketProjectGolang/pkg/repository"
	"BasketProjectGolang/pkg/router"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Println("Basket api starting...")

	// Set envs for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Set global logger
	logger.NewLogger(cfg)
	defer logger.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int64(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int64(time.Second)),
	}

	db := db.Connect(cfg)

	// Router group
	rootRouter := r.Group("")
	rootRouter.Use(mw.AuthMiddleware(cfg))
	//rootRouter.Use(mw.AuthenticateRole())
	apiRouter := r.Group(cfg.ServerConfig.RoutePrefix)

	// Repository
	productRepository := repository.NewProductRepository(db)
	if err != nil {
		log.Fatalf("ProductRepository error :%s", err.Error())
		return
	}
	authRepository := repository.NewAuthRepository(db)

	// Service
	productService := service.NewProductService(productRepository)
	authService := service.NewAuthService(authRepository)

	// Router
	productRouter := router.NewProductRouter(productService, authService)
	productRouter.Register(apiRouter)
	authRouter := router.NewAuthRouter(authService, *cfg)
	authRouter.Register(rootRouter)

	//Swagger
	if cfg.ServerConfig.Mode == "Development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.GET("", func(c *gin.Context) {
			c.Redirect(http.StatusSeeOther, "/swagger/index.html")
		})
	}

	// Health Check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	// Dependency Check
	r.GET("/readyz", func(c *gin.Context) {
		db, err := db.DB()
		if err != nil {
			zap.L().Fatal("cannot get sql database instance", zap.Error(err))
		}
		if err := db.Ping(); err != nil {
			zap.L().Fatal("cannot ping database", zap.Error(err))
		}
		c.JSON(http.StatusOK, nil)
	})

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Basket api has started")
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
