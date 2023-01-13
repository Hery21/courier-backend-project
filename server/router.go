package server

import (
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/handlers"
	" hery-ciaputra/final-project-backend/middlewares"
	" hery-ciaputra/final-project-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	RegisterService services.RegisterService
	AuthService     services.AuthService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Authorization",
		"Origin",
		"Accept",
		"X-Requested-With",
		"Content-Type",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}
	router.Use(cors.New(config))

	h := handlers.New(&handlers.HandlerConfig{
		RegisterService: c.RegisterService,
		AuthService:     c.AuthService,
	})

	router.POST("/register", middlewares.RequestValidator(func() any {
		return &dto.RegisterReq{}
	}), h.Register, middlewares.ErrorHandler)

	router.POST("/login", middlewares.RequestValidator(func() any {
		return &dto.LogInReq{}
	}), h.LogIn, middlewares.ErrorHandler)

	router.GET("/profile", middlewares.AuthorizeJWT, h.Profile, middlewares.ErrorHandler)

	router.PATCH("/profile", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.ProfileReq{}
	}), h.UpdateProfile, middlewares.ErrorHandler)

	router.PATCH("/top-up", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.TopUpReq{}
	}), h.TopUp, middlewares.ErrorHandler)

	router.GET("/addresses", middlewares.AuthorizeJWT, h.Address, middlewares.ErrorHandler)

	router.POST("/addresses", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.AddressReq{}
	}), h.CreateAddress, middlewares.ErrorHandler)

	router.PATCH("/addresses/:id", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.AddressReq{}
	}), h.EditAddress, middlewares.ErrorHandler)

	router.POST("/shippings", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.ShippingReq{}
	}), h.CreateShipping, middlewares.ErrorHandler)

	router.GET("/shippings", middlewares.AuthorizeJWT, h.ShippingList, middlewares.ErrorHandler)

	router.POST("/payment/:shipping-id", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.PaymentReq{}
	}), h.Payment, middlewares.ErrorHandler)

	router.POST("/promos", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.PromoReq{}
	}), h.AddPromo, middlewares.ErrorHandler)

	router.GET("/promos", middlewares.AuthorizeJWT, h.PromoList, middlewares.ErrorHandler)

	router.PATCH("/promos/:id", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.PromoReq{}
	}), h.EditPromo, middlewares.ErrorHandler)

	router.GET("/sizes", h.SizeList, middlewares.ErrorHandler)

	router.GET("/categories", h.CategoryList, middlewares.ErrorHandler)

	router.GET("/add-ons", h.AddOnList, middlewares.ErrorHandler)

	return router
}
