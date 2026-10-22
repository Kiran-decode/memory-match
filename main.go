package main

import (
	"log"
	"net/http"

	"memory-match/internal/game"
	"memory-match/internal/handlers"
	"memory-match/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := store.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("PostgreSQL connected successfully")

	gameService := game.NewService(db)
	gameHandler := handlers.NewGameHandler(gameService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "TEDx MindStage backend is running",
		})
	})

	router.Static("/assets", "web/assets")
	router.POST("/games", gameHandler.CreateGame)
	router.GET("/games/:id", gameHandler.GetGame)
	router.GET("/leaderboard", gameHandler.GetLeaderboard)
	router.POST("/games/:id/flip", gameHandler.FlipCard)
	router.POST("/games/:id/resolve", gameHandler.ResolveMismatch)
	router.GET("/", func(c *gin.Context) { c.File("web/index.html") })
	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
