package main

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/Vasu1712/qwiksift/server/internal/models"
    "github.com/Vasu1712/qwiksift/server/internal/scraper"
)

func main() {
    router := gin.Default()
    
    // API Endpoint
    router.GET("/api/products", func(c *gin.Context) {
        products, err := scraper.ScrapeAll()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch products",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "count":    len(products),
            "products": products,
        })
    })

    // Health Check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "ok",
            "version": "1.0.0",
            "time":    time.Now().Unix(),
        })
    })

    router.Run(":8080")
}
