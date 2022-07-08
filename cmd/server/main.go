package main

import (
	"log"
	"os"

	"github.com/adiletelf/payment-system-go/internal/handler"
	"github.com/adiletelf/payment-system-go/internal/middleware"
	"github.com/adiletelf/payment-system-go/internal/repository"
	"github.com/adiletelf/payment-system-go/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("API_SECRET") == "" {
		os.Setenv("API_SECRET", "defaultsecretkey")
	}
	os.Setenv("ACCESS_TOKEN_MINUTE_LIFESPAN", "15")
	os.Setenv("REFRESH_TOKEN_HOUR_LIFESPAN", "24")

	h := initializeBaseHandler()
	r := gin.Default()
	configureRoutes(r, h)
	r.Run(":8080")
}

func initializeBaseHandler() *handler.BaseHandler {
	db, err := util.SetupDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	tr := repository.NewTransactionRepo(db)
	ar := repository.NewAdminRepo(db)
	return handler.NewBaseHandler(tr, ar)
}

func configureRoutes(r *gin.Engine, h *handler.BaseHandler) {
	r.SetTrustedProxies([]string{"127.0.0.1"})

	public := r.Group("/api/admin")
	public.POST("/register", h.Register)
	public.POST("/login", h.Login)
	public.POST("/refresh", h.Refresh)

	protected := r.Group("/api")

	authenticationEnabled := os.Getenv("AUTHENTICATION_ENABLED")
	if authenticationEnabled == "true" {
		protected.Use(middleware.JwtAuthMiddleware())
	}

	protected.GET("/admin", h.CurrentAdmin)
	protected.GET("/transactions", h.GetAllTransactions)
	protected.POST("/transaction", h.CreateTransaction)
	protected.GET("/transaction/:id", h.GetTransactionStatus)
	protected.PUT("/transaction/:id", h.UpdateTransactionStatus)
	protected.GET("/transaction/cancel/:id", h.CancelTransaction)
}
