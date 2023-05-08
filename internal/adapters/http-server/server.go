package http_server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"auth-service/internal/adapters/http-server/handlers"
	"auth-service/internal/adapters/http-server/handlers/auth"
	"auth-service/internal/adapters/http-server/handlers/oauth"
	"auth-service/internal/adapters/http-server/session"
	"auth-service/internal/providers/service"
)

type Server struct {
	httpServer *http.Server
	log        *log.Logger
}

func NewServer(
	ctx context.Context,
	address string,
	provider service.Provider,
	log *log.Logger,
) (*Server, error) {
	publicURL, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	router := newRouter(log, provider, publicURL)

	httpServer := &http.Server{
		Addr:    publicURL.Host,
		Handler: router,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	return &Server{
		httpServer: httpServer,
		log:        log,
	}, nil
}

func newRouter(
	log *log.Logger,
	provider service.Provider,
	publicURL *url.URL,
) *gin.Engine {
	baseHandler := handlers.NewBaseHandler(log)
	authHandler := auth.NewAuthHandler(baseHandler, provider.GetAuthService())
	oauthHandler := oauth.NewOauthHandler(baseHandler, provider.GetOauthService())

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	sessionStore := cookie.NewStore([]byte("secret"))
	sessionStore.Options(sessions.Options{
		Path:     "/",
		Domain:   publicURL.Host,
		MaxAge:   60 * 60 * 24,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	router.Use(sessions.Sessions(session.DefaultSessionName, sessionStore))

	router.StaticFile("/auth/login", "./assets/login.html")
	router.StaticFile("/auth/confirm", "./assets/auth-confirm.html")
	router.StaticFile("/auth/error", "./assets/error.html")

	oauthGroup := router.Group("/oauth")
	{
		oauthGroup.GET("/authorize", oauthHandler.Authorize)
		oauthGroup.POST("/token", oauthHandler.Token)
		oauthGroup.POST("/validate", oauthHandler.Validate)
	}

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			authGroup := v1Group.Group("/auth")
			{
				authGroup.POST("/login", authHandler.Login)
				authGroup.GET("/user", authHandler.GetUser)
				authGroup.DELETE("/logout", authHandler.Logout)
				authGroup.POST("/register", authHandler.Register)
			}
		}
	}

	return router
}

func (s *Server) Start() error {
	s.log.Info(fmt.Sprintf("Server starting on addr: %v", s.httpServer.Addr))

	return s.httpServer.ListenAndServe()
}
