package main

// so common go cli commands
// go mod init github.com/username/repo
// go mod tidy
// go mod vendor
// go mod download
// go mod verify
// go mod graph
// go mod edit
// go mod why
// go mod why github.com/gin-gonic/gin
// build
// go build
// go build -o myapp
// go build -o myapp.exe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// building struct to define building data.
type building struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	EC       float64 `json:"ec"`
	OC       float64 `json:"oc"`
}

// buildings slice to seed record building data.
var buildings = []building{
	{ID: "1", Name: "Building 1", Location: "New York", EC: 100, OC: 200},
	{ID: "2", Name: "Building 2", Location: "San Francisco", EC: 150, OC: 250},
	{ID: "3", Name: "Building 3", Location: "Chicago", EC: 200, OC: 300},
}

// main function to boot up everything
func main() {
	router := gin.Default()
	router.GET("/buildings", getBuildings)
	router.GET("/buildings/:id", getBuildingByID)
	router.POST("/buildings", postBuildings)
	router.Run("0.0.0.0:8080")
}

// getBuildings responds with the list of all albums as JSON.
func getBuildings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, buildings)
}

// postbuildings adds an album from JSON received in the request body.
func postBuildings(c *gin.Context) {
	var newBuilding building

	// Call BindJSON to bind the received JSON to
	// newbuilding.
	if err := c.BindJSON(&newBuilding); err != nil {
		return
	}

	// Add the new building to the slice.
	buildings = append(buildings, newBuilding)
	c.IndentedJSON(http.StatusCreated, newBuilding)
}

// getbuildingByID locates the building whose ID value matches the id
// parameter sent by the client, then returns that building as a response.
func getBuildingByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of buildings, looking for
	// an building whose ID value matches the parameter.
	for _, b := range buildings {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "building not found"})
}
