package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func GetUsers(c *gin.Context) {
	type Users []User
	var users = Users{
		User{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		User{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}
	c.JSON(200, users)
	// curl -i http://localhost:8080/api/v1/users
}
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)
	if user_id == 1 {
		content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
		c.JSON(200, content)
	} else if user_id == 2 {
		content := gin.H{"id": user_id, "firstname": "Malcom", "lastname": "Merlyn"}
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
	}
	// curl -i http://localhost:8080/api/v1/users/1
}
func PostUser(c *gin.Context) {
	// The futur code…
}
func UpdateUser(c *gin.Context) {
	// The futur code…
}
func DeleteUser(c *gin.Context) {
	// The futur code…
}

