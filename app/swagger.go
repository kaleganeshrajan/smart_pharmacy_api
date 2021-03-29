package app

import (
	"smart_pharmacy_api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Swaggerinit function is to initialize swagger configurarion and swagger Link
func Swaggerinit(host string, port string) {
	docs.SwaggerInfo.Title = "Smart Pharmacy APIs"
	docs.SwaggerInfo.Description = "This is a Smart server."
	docs.SwaggerInfo.Host = host + ":" + port
	docs.SwaggerInfo.BasePath = ""

	url := ginSwagger.URL("http://" + docs.SwaggerInfo.Host + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
