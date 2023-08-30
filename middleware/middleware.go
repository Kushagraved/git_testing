package middleware

import "github.com/gin-gonic/gin"

//Request Middlware

// func Authenticate(c *gin.Context) {
// 	if(c.Request.Header.Get("Authorization") != "Bearer 1234"){
// 		c.JSON(401,gin.H{
// 			"error":"Not Authorized",
// 		})
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }

func Authenticate() gin.HandlerFunc {
	//Custom Logic

	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "Bearer 1234" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Not Authorized",
			})
			return
		}
		c.Next()
	}
}

//Response Middleware

func AddHeader() gin.HandlerFunc {
	//Custom Logic

	return func(c *gin.Context) {
		c.Writer.Header().Add("X-Developed-By", "Kushagra")
		c.Next()
	}
}
