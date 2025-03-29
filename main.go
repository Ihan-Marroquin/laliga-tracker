package main

import (
	"net/http"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Match struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	HomeTeam    string `json:"homeTeam" binding:"required"`
	AwayTeam    string `json:"awayTeam" binding:"required"`
	MatchDate   string `json:"matchDate" binding:"required"`
	Goals       int    `json:"goals"`
	YellowCards int    `json:"yellowCards"`
	RedCards    int    `json:"redCards"`
	ExtraTime   bool   `json:"extraTime"`
}

var db *gorm.DB

func main() {
	dsn := "host=" + getEnv("DB_HOST", "localhost") +
		" user=" + getEnv("DB_USER", "postgres") +
		" password=" + getEnv("DB_PASSWORD", "postgres") +
		" dbname=" + getEnv("DB_NAME", "laliga") +
		" port=" + getEnv("DB_PORT", "5432") +
		" sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos")
	}

	db.AutoMigrate(&Match{})

	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		matches := api.Group("/matches")
		{
			matches.GET("", getMatches)
			matches.GET("/:id", getMatch)
		}
	}

	router.Run(":8080")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getMatches(c *gin.Context) {
	var matches []Match
	db.Find(&matches)
	c.JSON(http.StatusOK, matches)
}

func getMatch(c *gin.Context) {
	id := c.Param("id")
	var match Match
	if result := db.First(&match, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.JSON(http.StatusOK, match)
}