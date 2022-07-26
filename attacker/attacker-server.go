package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Visible attack from client
	//r.StaticFile("/change-email", "./static/attacker-change-email.html")

	//Hidden attack
	r.StaticFile("/change-email", "./static/attacker-change-email_2.html")

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8090")
}
