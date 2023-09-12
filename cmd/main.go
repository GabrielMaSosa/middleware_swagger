package main

import (
	"os"

	"github.com/GabrielMaSosa/middleware-swagger/cmd/handlers"
	"github.com/GabrielMaSosa/middleware-swagger/cmd/middleware"
	"github.com/GabrielMaSosa/middleware-swagger/docs"
	product "github.com/GabrielMaSosa/middleware-swagger/internal/products"
	"github.com/GabrielMaSosa/middleware-swagger/pkg"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	gotenv.Load()

	// inyectamos las dependencias

	//tenemos que refactorizar el codigo viejo para agregar los
	//middleware

	slice, err := pkg.InitilizeBD("../products.json")
	if err != nil {
		panic(err)
	}

	repo := product.NewRepository(slice)
	servi := product.NewServiceProduct(&repo)
	hdler := handlers.NewHandlerProduct(servi)
	//fin de las inyecciones

	//inicio server
	server := gin.New()

	//vamos a agregar el middleware a todos los metodos pero si queremos separa por grupo
	//o rutas hacemos productsrout.Use(middleware)
	//importante el orden para el logger si ponemos por ultimo no atrapamos
	//los logs de los casos fallidos por no authoraization
	server.Use(
		gin.Recovery(),
		middleware.LoggerMiddleware(),
	)
	productsrout := server.Group("/products")
	productsrout.Use(middleware.AutorizationMiddleware())
	docs.SwaggerInfo.Host = os.Getenv("SERVER_ADDR")
	server.GET("/docs/*any", ginSwagger.WrapHandler(files.Handler))
	//vamos a agregar el middleware a todos los metodos pero si queremos separa por grupo
	//o rutas hacemos productsrout.Use(middleware) hacemos antes de definir las rutas
	handlers.Rutas(productsrout, hdler)

	server.Run(":8080")

}
