package main

import (
	"log"
	"net/http"
	"sg-stay-safe.org/restaurant/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("restaurant service...")
	router := gin.Default()
	router.GET("/restaurant/:id", getRestaurantById)

	router.Run(":5000")
}

// getRestaurantById locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getRestaurantById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	result := pkg.RetrieveById(id)

	c.IndentedJSON(http.StatusOK, result)
}
