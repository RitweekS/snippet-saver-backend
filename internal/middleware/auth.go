package middleware

import (
	"github.com/gin-gonic/gin"
)


func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve token directly from Authorization header
		// tokenString := c.Request.Header.Get("Authorization")
		// fmt.Println("Token String:", tokenString) // Debugging line to see token format

		// if tokenString == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		// 	c.Abort() // Stop further handlers from executing
		// 	return
		// }

		// // Parse and validate the JWT token
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	// Ensure token method is HMAC
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 	}
		// 	return jwtSecretKey, nil
		// })

		// if err != nil || !token.Valid {
		// 	fmt.Println("Token parse error:", err)
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		// 	c.Abort()
		// 	return
		// }

		// // Access claims if needed
		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	fmt.Println("Token claims:", claims)
		// 	// You can store claims in the context if needed
		// 	c.Set("claims", claims)
		// }

		// Proceed to the next handler
		c.Next()
	}
}
