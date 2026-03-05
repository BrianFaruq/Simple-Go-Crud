package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	initializers "github.com/GoProject/go-crud/Initializers"
	models "github.com/GoProject/go-crud/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// CHANGE: If the cookie is missing/empty, you MUST return after aborting.
	// What's wrong before: AbortWithStatus was called, but the function continued,
	// leading to parsing an empty token and/or hitting DB calls.
	tokenString, err := c.Cookie("Authorization")
	if err != nil || tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// CHANGE: Never use log.Fatal in request middleware.
	// What's wrong before: log.Fatal exits the entire server process on any JWT parse error.
	// What should happen: treat it as an auth failure (401) and continue serving other requests.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil || token == nil || !token.Valid {
		log.Println("jwt parse/validate error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// CHANGE: Type-safe exp handling.
		// What's wrong before: claims["exp"].(float64) will panic if exp is missing or not a float64.
		// What should happen: if exp is missing/invalid, treat as unauthorized.
		exp, ok := claims["exp"].(float64)
		if !ok || float64(time.Now().Unix()) > exp {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//Find user with the token
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			// CHANGE: Must return after abort.
			// What's wrong before: code continued and could attach an empty user / continue the chain.
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//Attach to request
		c.Set("user", user)

		// CHANGE: Removed fmt.Println debug of non-existent claims.
		// What's wrong before: you never set "foo"/"nbf" claims, so it adds noise and imports fmt.
		c.Next()

	} else {
		// CHANGE: Must return after abort.
		// What's wrong before: no return; function could continue in future edits.
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
