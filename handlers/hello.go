package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func HelloHandler(c *gin.Context) {
    c.String(http.StatusOK, "Hello, World!")
}

// GreetHandler godoc
// @Summary Greet user by name
// @Description greet user with provided name
// @Tags hello
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Success 200 {object} map[string]string
// @Router /greet/{name} [get]
func GreetHandler(c *gin.Context) {
    name := c.Params.ByName("name")
    c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
}

