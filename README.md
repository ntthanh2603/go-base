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
	return Controller("/hello-world",
		Get("/",
		   func() any {
		   	return appService.HelloWorldGet()
		   },
		   UseGuard(TestGuard),
		),
	)
}
```

## Service
```go
package services

type AppServiceType struct {
}

func AppService() *AppServiceType {
	return &AppServiceType{}
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func (AppService *AppServiceType) HelloWorldGet() HelloWorldResponse {
	return HelloWorldResponse{
		Message: "[GET] Hello World",
	}
}
```

## Env
```
POSTGRES_DSN="host=<replace your host> user=<replace your username> password=<replace your password> dbname=<replace your database name> port=<replace your database port>" 
```
