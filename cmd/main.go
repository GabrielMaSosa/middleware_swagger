package main

import (
	"fmt"
	"os"

	"github.com/GabrielMaSosa/middleware-swagger/cmd/handlers"
	"github.com/GabrielMaSosa/middleware-swagger/cmd/middleware"
	product "github.com/GabrielMaSosa/middleware-swagger/internal/products"
	"github.com/GabrielMaSosa/middleware-swagger/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("TOKEN", "123456")
	path := os.Getenv("MYPATH")
	// inyectamos las dependencias
	fmt.Println("mivar", path)

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
		middleware.AutorizationMiddleware(),
	)
	productsrout := server.Group("/products")
	//vamos a agregar el middleware a todos los metodos pero si queremos separa por grupo
	//o rutas hacemos productsrout.Use(middleware) hacemos antes de definir las rutas
	handlers.Rutas(productsrout, hdler)

	server.Run(":8080")

}
