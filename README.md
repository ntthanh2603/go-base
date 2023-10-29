**🧡Go Base is a [go-based](https://gin-gonic.com/) project that provides a quick way to build Rest APIs.**

# ⚡️ Quickstart

## App

```go
func App() {
	serverConfig := ServerConfig{

		Controllers: []Controller{
			controllers.MyController,
		},

		Middlewares: []gin.HandlerFunc{
			middlewares.AuthGuard(),
		},

		Port: ":3000",

		DebugLogger: true, // optional
	}

	CreateServer(serverConfig)
}

```

## Controller

```go
func MyController(r *gin.RouterGroup) {
    myController := r.Group("/my-controller")
    myService := services.MyService()
    myController.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

    myController.POST("/",  myService.HelloWorld)
}
```

## Service

```go
type MyServiceType struct {
}

func MyService() *myServiceType {
	return &myServiceType{}
}

func (myService *MyServiceType) HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
}
```
