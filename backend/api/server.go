package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/indramahaarta/chitchat/db/sqlc"
	"github.com/indramahaarta/chitchat/docs"
	"github.com/indramahaarta/chitchat/token"
	"github.com/indramahaarta/chitchat/util"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{config: *config, store: store, tokenMaker: tokenMaker}
	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// configure swagger docs
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = server.config.BackendSwaggerHost
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// health check api
	router.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})

	// auth
	router.POST("/api/auth/signin", server.signin)
	router.POST("/api/auth/signup", server.signup)
	router.POST("/api/auth/google", server.google)

	// social
	authRouter.GET("/api/users", server.getSocialUser)
	// authRouter.GET("/api/users/:uid", server.getUserDetails)

	// chats
	router.GET("/ws/chats", server.chat)
	authRouter.GET("/api/chats/histories", server.getChatsHistories)
	authRouter.GET("/api/chats/:uid", server.getChatsDetails)

	server.router = router
}

func getUserPayload(ctx *gin.Context) (*token.Payload, error) {
	payload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		return nil, fmt.Errorf("payload is missing")
	}
	userPayload, ok := payload.(*token.Payload)
	if !ok {
		return nil, fmt.Errorf("payload structure is not corrent")
	}

	return userPayload, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
