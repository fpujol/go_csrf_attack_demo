package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Email string `form:"email"`
	Password string `form:"password"`
}

type ChangeEmail struct {
	Email string `form:"email"`
}


func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("my-secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Static("/assets", "./assets")
	r.StaticFile("/login", "./static/login.html")
	r.StaticFile("/invalid-login", "./static/invalid-login.html")
	
	r.POST("/login", login)

	authorized := r.Group("/", AuthRequired())
	authorized.Use(AuthRequired())
	{
		authorized.StaticFile("/profile", "./static/profile.html")
		authorized.StaticFile("/email-changed", "./static/email-changed.html")
		r.POST("/change-email", changeEmail)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func login(c *gin.Context) {
	var l Login
    c.Bind(&l)

	if l.Email=="your-email@mail.com" && l.Password=="******" {
		session := sessions.Default(c)
		session.Set("user", l.Email)
      	session.Save()
		c.Redirect(http.StatusFound, "/profile")
	} else {
		c.Redirect(http.StatusFound, "/invalid-login")
	}
}

func changeEmail(c *gin.Context) {	
	log.Println("Email Changed")
	c.Redirect(http.StatusFound, "/email-changed")
}


func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		user := session.Get("user")
		log.Println("user: ",user)
		if user ==nil {
			c.Redirect(http.StatusFound, "/login")
			return			
		}

		// before request
		c.Next()

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
