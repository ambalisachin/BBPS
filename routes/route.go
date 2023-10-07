package routes

import (
	"BBT/controllers"
	middleware "BBT/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter func  sets up a router using the gin web framework.
func SetupRouter() *gin.Engine {
	//	creates a router as 'r' & sets it to use the gin framework's default settings.
	//This allows the router to use all of the default routes and middleware functions that are available in the gin framework.
	r := gin.Default()
	//creates a new router group named 'v1' which is associated with the URL path prefix '/v1'. This router group can be used to handle routes specific to the '/v1' prefix.
	v1 := r.Group("/v1")
	//creates a route group  called "v1". The route group is called "/add" and can contain routes that are related to adding something.
	add := v1.Group("/add")
	//using a middleware called DecryptRequest().
	//This middleware is responsible for decrypting requests that are sent to the server.
	// It will decrypt any encrypted requests that are made to the server, allowing the server to access any data that is encrypted.
	// The middleware will also ensure that all requests are correctly authenticated and authorized, preventing unauthorized access to sensitive data.
	add.Use(middleware.DecryptRequest())
	{
		add.POST("/biller", controllers.CreateABiller)

		add.POST("/prepaid", controllers.PrepaidAPI)

		add.POST("/billercategoty", controllers.BillerCategory)

		add.POST("/fetchbill", controllers.NewHandlerDb().FetchBill)

		add.POST("/billvalidation", controllers.NewHandlerDb().BillValidation)

		add.POST("/nonprepaidbill", controllers.NewHandlerDb().NonPrepaid)

		v1.GET("billercategory", controllers.GetBillerCategory)

		v1.GET("billercategory/:id", controllers.BillerCategoryById)

		v1.GET("biller/:id", controllers.BillFetch)

		v1.POST("/token", controllers.Login)

		v1.POST("/validate", controllers.ValidateABill)

		secured := v1.Group("/secured")
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	//creates a route group called "encrypt" with the path of "/data". Any routes within the encrypt route group will have the path prefix of "/data".
	encrypt := r.Group("/data")
	{
		encrypt.POST("encrypt", controllers.EncryptDataHandler)
		encrypt.POST("decrypt", controllers.DecryptDataHandler)
	}
	return r
}
