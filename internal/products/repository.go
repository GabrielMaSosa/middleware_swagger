package product

import "github.com/GabrielMaSosa/middleware-swagger/internal/domain"

type ProductRepository interface {
	GetAll() (ret []domain.Product, err error)
	Update(id int, data domain.Product) (ret *domain.Product, err error)
	Delete(id int) (ret *domain.Product, err error)
	GetById(id int) (ret *domain.Product, err error)
	PartialUpdate(id int, data map[string]interface{}) (ret *domain.Product, err error)
	GetPriceMayor(price float64) (dt []domain.Product, err error)
	SaveNewProduct(data domain.Product) (ret *domain.Product, err error)
}

// esta estructura vamos a usar solo para el PATCH
type DataProductRepository struct {
	ID           int     `val:"id"`
	Name         string  `val:"name"`
	Quantity     int     `val:"quantity"`
	Code_value   string  `val:"code_value"`
	Is_published bool    `val:"is_published"`
	Expiration   string  `val:"expiration"`
	Price        float64 `val:"price"`
}
