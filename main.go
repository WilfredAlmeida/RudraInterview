package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(CORSMiddleware())
	err := router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "Hello Heckerr",
			"payload": nil,
		})
	})

	router.POST("/v1/getRules", GetRulesApiHandler)

	router.POST("/v1/addRules", AddRulesApiHandler)

	router.POST("/v1/updateRuleStatus", UpdateStatusApiHandler)

	err = router.Run(":6565")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func AddRulesApiHandler(c *gin.Context) {

	c.IndentedJSON(http.StatusCreated, gin.H{
		"status":  1,
		"message": "Rules Created",
		"payload": nil,
	})

}

func GetRulesApiHandler(c *gin.Context) {

	c.IndentedJSON(http.StatusCreated, gin.H{
		"status":  1,
		"message": "Rules Fetched Successfully",
		"payload": gin.H{
			"rules": "SomeRules",
		},
	})

}

func UpdateStatusApiHandler(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
