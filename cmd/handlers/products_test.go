package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
	product "github.com/GabrielMaSosa/middleware-swagger/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	os.Setenv("TOKEN", "123456")
	var slice = []domain.Product{
		{
			ID:           1,
			Name:         "Oil - Margarine",
			Quantity:     439,
			Code_value:   "S82254D",
			Is_published: true,
			Expiration:   "15/12/2021",
			Price:        71.42,
		},
		{
			ID:           2,
			Name:         "Pineapple - Canned, Rings",
			Quantity:     345,
			Code_value:   "M4637",
			Is_published: true,
			Expiration:   "09/08/2021",
			Price:        352.79,
		},
	}

	repo := product.NewRepositoryTesting(slice)
	servi := product.NewServiceProduct(&repo)
	hdler := NewHandlerProduct(servi)
	//inicio server
	server := gin.New()
	productsrout := server.Group("/products")

	Rutas(productsrout, hdler)
	return server
}

func TestFuncionalHandlerProductGetByID(t *testing.T) {

	t.Run("Happy path get by id", func(t *testing.T) {
		//arrange
		//levanto servrer
		expectedStatusCode := 200
		expectBodty := `{"id":1,"name":"Oil - Margarine","quantity":439,"code_value":"S82254D","is_published":true,"expiration":"15/12/2021","price": 71.42}`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		//--> input
		request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		//act
		r.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, string(expectBodty), response.Body.String(), "error")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})
	t.Run("Happy path get all", func(t *testing.T) {
		//arrange
		expectedStatusCode := 200
		expectBodty := `[{"id": 1,"name": "Oil - Margarine","quantity": 439,"code_value": "S82254D","is_published": true,"expiration": "15/12/2021","price": 71.42},
		{"id": 2,"name": "Pineapple - Canned, Rings","quantity": 345,"code_value": "M4637","is_published": true,"expiration": "09/08/2021","price": 352.79}]`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// --> input
		request := httptest.NewRequest(http.MethodGet, "/products", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		// act
		r.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, string(expectBodty), response.Body.String(), "error")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})
	t.Run("Happy path POST", func(t *testing.T) {
		//arrange
		expectedStatusCode := 201
		var expectBodty = "null"

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// --> input
		mibdy := domain.Product{
			ID:           8,
			Name:         "aceite - Margarine",
			Quantity:     439,
			Code_value:   "S82254DA",
			Is_published: true,
			Expiration:   "15/12/2025",
			Price:        71.42,
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mibdy)

		request := httptest.NewRequest(http.MethodPost, "/products/add", &buf)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		// act
		r.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectBodty, response.Body.String(), "error POST")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})

	t.Run("Happy path DELETE", func(t *testing.T) {
		//arrange
		expectedStatusCode := 200
		var expectBodty = `{"msg": "Succesful delete"}`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		request := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		// act
		r.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectBodty, response.Body.String(), "error DELETE")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})

	//falta correr el resto de los casos con errores
	t.Run("One error id negative get by id", func(t *testing.T) {
		//arrange
		//levanto servrer
		expectedStatusCode := http.StatusBadRequest
		expectBodty := `{"mesage": "Bad request2"}`

		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		//--> input
		request := httptest.NewRequest(http.MethodGet, "/products/-1", nil)
		request.Header.Add("token", "123456")
		response := httptest.NewRecorder()
		r := createServer()
		//act
		r.ServeHTTP(response, request)

		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, string(expectBodty), response.Body.String(), "error")
		assert.Equal(t, expectedHeaders, response.Header())
		//assert

	})

}
