package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
//
// Struct tags such as json:"artist" specify what a field’s name should be
// when the struct’s contents are serialized into JSON.
// Without them, the JSON would use the struct’s capitalized field names
// – a style not as common in JSON.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Gin enginge initialization.
	app := gin.Default()

	// Routing.
	app.GET("/albums", getAlbums)
	app.GET("/albums/:id", getAlbumByID)
	app.POST("/albums", postAlbums)

	// REST_Requests.http file info.
	fmt.Println("############################################################################")
	fmt.Println("Check the REST_Requests.http file, there are some examples of HTTP requests.")
	fmt.Println("############################################################################")

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	// Note: this method will block the calling goroutine indefinitely unless an error happens.
	app.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
//
// Write a getAlbums function that takes a gin.Context parameter.
// Note that you could have given this function any name – neither Gin nor Go
// require a particular function name format.
//
// gin.Context is the most important part of Gin. It carries request details,
// validates and serializes JSON, and more.
// (Despite the similar name, this is different from Go’s built-in context package.)
func getAlbums(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client.
	// Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON reveived in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	// Add a 201 status code to the response, along with JSON representing the album you added.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client,
// then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// 	Param from path /albums/:id
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
