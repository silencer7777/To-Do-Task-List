// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type User struct {
    Password string
    Tasks    []Task
}
type Task struct {
    Text      string `json:"text"`
    Completed bool   `json:"completed"`
}

var users = make(map[string]User)

func main() {
    r := gin.Default()
    r.POST("/register", func(c *gin.Context) {
        var req struct{ Username, Password string }
        c.BindJSON(&req)
        if _, exists := users[req.Username]; exists {
            c.JSON(http.StatusBadRequest, gin.H{"error": "User exists"})
            return
        }
        users[req.Username] = User{Password: req.Password, Tasks: []Task{}}
        c.JSON(http.StatusOK, gin.H{"message": "Registered"})
    })
    r.POST("/login", func(c *gin.Context) {
        var req struct{ Username, Password string }
        c.BindJSON(&req)
        user, exists := users[req.Username]
        if !exists || user.Password != req.Password {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
    })
    r.GET("/tasks/:username", func(c *gin.Context) {
        username := c.Param("username")
        user, exists := users[username]
        if !exists {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, user.Tasks)
    })
    r.POST("/tasks/:username", func(c *gin.Context) {
        username := c.Param("username")
        var task Task
        c.BindJSON(&task)
        user := users[username]
        user.Tasks = append(user.Tasks, task)
        users[username] = user
        c.JSON(http.StatusOK, user.Tasks)
    })
    r.Run(":8080")
}
