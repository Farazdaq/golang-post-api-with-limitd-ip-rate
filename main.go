// This file mange the appliction router
package main

import (
	"POST-API/controllers"
	"fmt"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	// "errors"
)

// This function listen the client IP
func keyFunc(c *gin.Context) string {
	fmt.Print(c.ClientIP())
	return c.ClientIP()
}

// This function handels the error response when the post rate is out of limit
func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too Many Post, try later"+time.Until(info.ResetTime).String())
}

func main() {
	router := gin.Default()
	// the line identfy the rate limit and time the clients post requst
	store := ratelimit.RedisStore(&ratelimit.RedisOptions{RedisClient: redis.NewClient(&redis.Options{Addr: "loclhost:7680"}),
		Rate:  time.Second,
		Limit: 10,
	})
	// This function evalute rate limit
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{ErrorHandler: errorHandler, KeyFunc: keyFunc})

	// This the post route listen on port 8080 the post and call the search function in contorller
	router.POST("/country", mw, controllers.CearchCuntry)
	router.Run("localhost:8080")

}
