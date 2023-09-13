package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
	product "github.com/GabrielMaSosa/middleware-swagger/internal/products"
	"github.com/GabrielMaSosa/middleware-swagger/pkg/web"
	"github.com/gin-gonic/gin"
)

// agrego la interfaz para la composicion
type HandlerProduct struct {
	service product.ProductService
}

// Constructor del Handler
func NewHandlerProduct(sv *product.ProductServiceImpl) *HandlerProduct {

	return &HandlerProduct{service: sv}
}

// Definicion de las rutas
func Rutas(g *gin.RouterGroup, h *HandlerProduct) {
	g.GET("", h.GetAll())
	g.PUT("/:id", h.Update())
	g.DELETE("/:id", h.Delete())
	g.PATCH("/:id", h.PartialSave())
	g.GET("/:id", h.GetProductByIdPatch())
	g.GET("/search", h.SearchProduct())
	g.POST("/add", h.Save())
}

var (
	ErrNotFound = errors.New("Not found")
)

func (h *HandlerProduct) PartialSave() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}
		if idn <= 0 {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}

		dta := map[string]interface{}{}
		if err := ctx.ShouldBindJSON(&dta); err != nil {
			web.RequestError(ctx, err.Error(), http.StatusBadRequest)
			return
		}
		if err2 := ValidateRequest(dta); err2 != nil {
			web.RequestError(ctx, err2.Error(), http.StatusBadRequest)
			fmt.Println(err2)
			return
		}

		dta1, err5 := h.service.UpdatePartial(idn, dta)
		if err5 != nil {
			web.RequestError(ctx, err5.Error(), http.StatusInternalServerError)
			return
		} else {
			web.Requestok(ctx, 200, dta1)
			return
		}

	}
}

// ListProducts godoc
// @Summary Listar los productos
// @Tags Products
// @Description Obtienes un slice de productos
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} []domain.Product
// @Failure 404 {object} string
// @Router /products [get]
func (h *HandlerProduct) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dta, err := h.service.ServiceGetAll()
		if err != nil {
			web.RequestError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
		web.Requestok(ctx, 200, dta)
		return
	}
}

func (h *HandlerProduct) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var datain domain.Product

		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}
		if idn <= 0 {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}
		if err2 := ctx.ShouldBindJSON(&datain); err2 != nil {
			web.RequestError(ctx, err2.Error(), http.StatusBadRequest)
			return

		}
		if err3 := ValidateData(datain); err3 != nil {
			web.RequestError(ctx, err3.Error(), http.StatusBadRequest)
			return
		}
		val2, err5 := h.service.UpdateItem(idn, datain)
		if err5 != nil {
			web.RequestError(ctx, err5.Error(), http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, val2)
		return
	}
}

func (h *HandlerProduct) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}
		if idn <= 0 {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}

		_, err5 := h.service.Delete(idn)
		if err5 != nil {
			web.RequestError(ctx, err5.Error(), http.StatusNotFound)
			return
		}
		web.Requestok(ctx, 200, "Delete")
		return
	}

}

// GetProduct traer un producto por id
func (h *HandlerProduct) GetProductByIdPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}
		if idn <= 0 {
			web.RequestError(ctx, ErrAtributenovalid.Error(), http.StatusBadRequest)
			return
		}

		valore, err := h.service.ServiGetById(idn)
		if err != nil {
			web.RequestError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		web.Requestok(ctx, 200, valore)
		return
	}
}

// SearchProduct traer un producto por nombre o categoria
func (h *HandlerProduct) SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		query := ctx.Query("priceGt")
		priceGt, err := strconv.ParseFloat(query, 32)
		if err != nil {

			web.RequestError(ctx, ErrparamPrice.Error(), http.StatusBadRequest)
			return
		}
		valore, err := h.service.ServiGetPriceMayor(priceGt)
		if err != nil {
			web.RequestError(ctx, fmt.Sprintf("%g", priceGt), http.StatusNotFound)
			return
		}

		web.Requestok(ctx, 200, valore)
		return
	}
}

func (h *HandlerProduct) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dta := domain.Product{}
		if err := ctx.ShouldBindJSON(&dta); err != nil {

			web.RequestError(ctx, err.Error(), http.StatusBadRequest)

			return
		}

		if _, err1 := ValidateEmpty(&dta); err1 != nil {

			web.RequestError(ctx, err1.Error(), http.StatusBadRequest)
			return
		}

		if err2 := ValidateCodeValue(&dta); err2 != nil {

			web.RequestError(ctx, err2.Error(), http.StatusBadRequest)
			return

		}
		if err3 := ValidateDate(&dta); err3 != nil {
			web.RequestError(ctx, err3.Error(), http.StatusBadRequest)
			return

		}
		//esto hago para agregar a mi db simulada
		valxx, errxx := h.service.ServiNewItem(dta)
		if errxx != nil {
			web.RequestError(ctx, errxx.Error(), http.StatusInternalServerError)
			return
		}

		web.Requestok(ctx, 200, valxx)
		return
	}
}
