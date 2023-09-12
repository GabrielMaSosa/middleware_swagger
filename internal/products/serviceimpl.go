package product

import "github.com/GabrielMaSosa/middleware-swagger/internal/domain"

type ProductServiceImpl struct {
	//compongo con la iterfaz del repositorio
	repository ProductRepository
}

// Constructor Service
func NewServiceProduct(repo ProductRepository) *ProductServiceImpl {

	return &ProductServiceImpl{repository: repo}

}

func (s *ProductServiceImpl) ServiceGetAll() (dt []domain.Product, err error) {
	dt, err = s.repository.GetAll()
	return

}

// hace update para este caso pasa mano ya que no hay mucho que hacer
// pero sirve para manejar el mod de trabajo
func (s *ProductServiceImpl) UpdateItem(id int, data domain.Product) (dt *domain.Product, err error) {

	dt, err = s.repository.Update(id, data)

	return

}
func (s *ProductServiceImpl) Delete(id int) (dt *domain.Product, err error) {
	dt, err = s.repository.Delete(id)
	return
}

func (s *ProductServiceImpl) UpdatePartial(id int, data map[string]interface{}) (dt *domain.Product, err error) {

	dt, err = s.repository.PartialUpdate(id, data)
	return

}

func (s *ProductServiceImpl) ServiGetById(id int) (dt *domain.Product, err error) {
	dt, err = s.repository.GetById(id)
	return

}
func (s *ProductServiceImpl) ServiGetPriceMayor(price float64) (dt []domain.Product, err error) {
	dt, err = s.repository.GetPriceMayor(price)
	return
}

func (s *ProductServiceImpl) ServiNewItem(data domain.Product) (ret *domain.Product, err error) {
	ret, err = s.repository.SaveNewProduct(data)
	return
}
