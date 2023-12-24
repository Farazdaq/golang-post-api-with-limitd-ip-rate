// Ths the main Conterllers that manage the app logic
package controllers

// the part to import the pack and laibreis need
import (
	"POST-API/models"

	"github.com/gin-gonic/gin"
)

// This function is lunched by the post clint requst

// the call the data model then struct then search for the country name and printed out
func CearchCuntry(c *gin.Context) {

	// the varable holed the user posted value
	name := c.DefaultPostForm("name", "Afghanistan")

	//This loop to go accros the struct to fin the country info
	for i, cunt := range models.Counts {
		if cunt.Name == name {
			c.String(200, "info %v", &models.Counts[i])
		}
	}

}
