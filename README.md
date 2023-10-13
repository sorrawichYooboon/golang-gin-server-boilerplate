# How to set up golang gin server boiler plate

1. Create a new directory for your project and navigate to it in your terminal.
2. Run the following command to initialize a new Go module
   `go mod init <module-name>`
   Replace <module-name> with the name of your module.
3. Run the following command to install the Gin framework
   `go get -u github.com/gin-gonic/gin`
4. Create a new file called main.go in the root of your project and add the following code

```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a route for the ping API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	// Start the server
	router.Run(":3000")
}
```

5. Save the main.go file and run the following command to start the server:
   `go run main.go`
