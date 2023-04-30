package http_server

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internal/handlers/http-server/auth"
	"auth-service/internal/handlers/http-server/oauth"
	"auth-service/internal/providers/service"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(
	ctx context.Context,
	address string,
	provider service.Provider,
) (*Server, error) {
	router := newRouter(provider)

	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	return &Server{
		httpServer: httpServer,
	}, nil
}

func newRouter(provider service.Provider) *gin.Engine {
	authHandler := auth.NewAuthHandler(provider.GetAuthService())
	oauthHandler := oauth.NewOauthHandler(provider.GetOauthService())

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/auth/login", "./assets/login.html")
	router.StaticFile("/auth/confirm", "./assets/auth-confirm.html")

	oauthGroup := router.Group("/oauth")
	{
		oauthGroup.GET("/authorize", oauthHandler.Authorize)
		oauthGroup.GET("/token", oauthHandler.Token)
		oauthGroup.POST("/validate", oauthHandler.Validate)
	}

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			authGroup := v1Group.Group("/auth")
			{
				authGroup.POST("/login", authHandler.Login)
			}
		}
	}

	return router
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
