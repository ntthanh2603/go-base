**🧡Go Base is a [Gin-based](https://gin-gonic.com/) project that provides a quick way to build Rest APIs.**
    
# ⚡️ Quickstart

## App
```go
func App() *gin.Engine {
	// Initialize the environment
	env.Init()

	// Connect to the database
	DatabaseConnect()
	// Define the server configuration
	serverConfig := ServerConfig{

		Controllers: []Controller{
			controllers.AppController()
		},

		Middlewares: []gin.HandlerFunc{
			middlewares.Cors(),
			middlewares.Custom(),
		},

		Port: configs.Port,

		DebugLogger: true,
	}

	// Create the server
	return CreateServer(serverConfig)

}

```
## Controller
```go
func AppController() *gin.Engine {
	appService := services.AppService()
	appController := ControllerBase{
		basePath: "/hello-world",
		routes: []RouteBase{
			Get("/", appService.HelloWorldGet),
			Post("/", appService.HelloWorldPost),
		},
	}
	return Controller(appController)
}
```

## Service
```go
func AppService() *AppServiceType {
	return &AppServiceType{}
}

func (AppService *AppServiceType) HelloWorldPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "[POST] Hello World",
	})
}

func (AppService *AppServiceType) HelloWorldGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "[GET] Hello World",
	})
}

```

## Env
```
POSTGRES_DSN="host=<replace your host> user=<replace your username> password=<replace your password> dbname=<replace your database name> port=<replace your database port>" 
```
