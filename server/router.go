package server

import (
    "github.com/gin-gonic/gin"
    "topcoder.com/skill-builder/golang/controllers"
)

func NewRouter() *gin.Engine {
    router := gin.New()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    health := new(controllers.HealthController)

    router.GET("/health", health.Status)

    v1 := router.Group("/api/v1")
    {
        recoveryRoutes := new(controllers.RecoveryController)
        v1.POST("/recover", recoveryRoutes.RecoverData)

        //TODO: Add appropriate routes and controllers for Medium-500 & Hard-1000
    }

    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{
            "message": "Not Found",
        })
    })

    return router
}
