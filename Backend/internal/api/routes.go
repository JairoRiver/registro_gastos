package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter(frontendUrl string) {
	router := gin.Default()

	// Allow CORS policy
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{frontendUrl}
	router.Use(cors.New(configCors))

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	//login routes
	router.POST("/v1/user/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	//user routes
	router.POST("/v1/user", server.createUser)
	authRoutes.GET("/v1/user/:id", userMiddleware(server), server.getUser)
	authRoutes.PUT("/v1/user/:id", userMiddleware(server), server.updateUser)
	authRoutes.DELETE("/v1/user/:id", userMiddleware(server), server.deleteUser)

	//group routes
	authRoutes.POST("/v1/group", userMiddleware(server), server.createGroup)
	authRoutes.GET("/v1/group/:id", userMiddleware(server), server.getGroup)
	authRoutes.GET("/v1/groups", userMiddleware(server), server.listGroups)
	authRoutes.PUT("/v1/group/:id", userMiddleware(server), server.updateGroup)
	authRoutes.DELETE("/v1/group/:id", userMiddleware(server), server.deleteGroup)

	//type routes
	router.POST("/v1/type", server.createType)
	router.GET("/v1/type/:id", server.getType)
	router.GET("/v1/types", server.listTypes)
	router.PUT("/v1/type/:id", server.updateType)
	router.DELETE("/v1/type/:id", server.deleteType)

	//entry routes
	authRoutes.POST("/v1/entry", server.createEntry)
	authRoutes.GET("/v1/entry/user/:userid/entry/:entryid", userMiddleware(server), server.getEntry)
	authRoutes.GET("/v1/entries/:id", userMiddleware(server), server.listEntries)
	authRoutes.PUT("/v1/entry/:id", userMiddleware(server), server.updateEntry)
	authRoutes.DELETE("/v1/entry/:id", userMiddleware(server), server.deleteEntry)

	server.router = router
}
