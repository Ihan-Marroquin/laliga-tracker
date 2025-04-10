	package main

	import (
		"net/http"
		"os"
		"github.com/gin-contrib/cors"
		"github.com/gin-gonic/gin"
		"gorm.io/driver/postgres"
		"gorm.io/gorm"
		_ "laliga-tracker/docs"
		swaggerFiles "github.com/swaggo/files"
		ginSwagger "github.com/swaggo/gin-swagger"
	)

	// @title La Liga Tracker API
	// @version 1.0
	// @description API para gestionar partidos de La Liga
	// @host localhost:8080
	// @BasePath /api

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

	// @Summary Listar todos los partidos
	// @GET /matches
	// @Produce json
	// @Success 200 {array} Match

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
				matches.POST("", createMatch)
				matches.PUT("/:id", updateMatch)
				matches.DELETE("/:id", deleteMatch)
				
				matches.PATCH("/:id/goals", incrementGoals)
				matches.PATCH("/:id/yellowcards", incrementYellowCards)
				matches.PATCH("/:id/redcards", incrementRedCards)
				matches.PATCH("/:id/extratime", setExtraTime)
			}
		}

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	func createMatch(c *gin.Context) {
		var newMatch Match
		if err := c.ShouldBindJSON(&newMatch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		db.Create(&newMatch)
		c.JSON(http.StatusCreated, newMatch)
	}

	func updateMatch(c *gin.Context) {
		id := c.Param("id")
		var match Match
		if result := db.First(&match, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
			return
		}

		var updateData Match
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&match).Updates(updateData)
		c.JSON(http.StatusOK, match)
	}

	func deleteMatch(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&Match{}, id)
		c.Status(http.StatusNoContent)
	}

	func updateCounter(c *gin.Context, field string) {
		id := c.Param("id")
		var match Match
		if result := db.First(&match, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
			return
		}

		db.Model(&match).Update(field, gorm.Expr(field + " + ?", 1))
		c.JSON(http.StatusOK, match)
	}

	func incrementGoals(c *gin.Context) {
		updateCounter(c, "goals")
	}

	func incrementYellowCards(c *gin.Context) {
		updateCounter(c, "yellow_cards")
	}

	func incrementRedCards(c *gin.Context) {
		updateCounter(c, "red_cards")
	}

	func setExtraTime(c *gin.Context) {
		id := c.Param("id")
		var match Match
		if result := db.First(&match, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
			return
		}

		db.Model(&match).Update("extra_time", true)
		c.JSON(http.StatusOK, match)
	}

