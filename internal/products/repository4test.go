package product

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
)

type RepositoryImplTesting struct {
	bd []domain.Product
}

//estructura con puntero hacia la base de datos este caso slice exportados

//constructor

func NewRepositoryTesting(dbex []domain.Product) (ret RepositoryImplTesting) {
	ret.bd = dbex
	return
}

func (r *RepositoryImplTesting) GetAll() (ret []domain.Product, err error) {
	ret = r.bd
	return

}

func (r *RepositoryImplTesting) Update(id int, data domain.Product) (ret *domain.Product, err error) {
	fmt.Println("status repo", data)
	fmt.Println(id)

	flag := false
	for i, v := range r.bd {
		if v.ID == id {
			//atento a las asignaciones con puntero
			//repasar mas
			v = data
			ret = &data
			flag = true
			r.bd[i] = data
			fmt.Println("valor despues del cambio", v)
			fmt.Println("encontre repo")
		}
	}

	if flag == false {
		//el item no esta enonces debemos agregarlo

		r.bd = append(r.bd, data)

		err = nil
		ret = &data
		return
	}
	if flag {

	}
	//no hay muchos errores que poner aca
	//ya que si llegamos hasta aca
	//es porque ya esta todo validado
	err = nil
	return
}

func (r *RepositoryImplTesting) Delete(id int) (ret *domain.Product, err error) {
	//se puede dar que el id del producto no sea
	//el mismo que el id de la ubicacion por lo tanto
	//hay que separa estos valores

	flag := false
	index := 0
	for i, v := range r.bd {
		if v.ID == id {
			index = i
			fmt.Println("Encontrado")
			//con el append de abajo borramo
			r.bd = append(r.bd[:index], r.bd[index+1:]...)

			flag = true
		}
	}
	if flag {
		err = nil
	}

	return
}

func (r *RepositoryImplTesting) PartialUpdate(id int, data map[string]interface{}) (ret *domain.Product, err error) {
	fmt.Println("---------------------")
	fmt.Println("-----------REPO----------")
	flag := false

	indice := 0
	src := DataProductRepository{}
	for i, v := range r.bd {
		if v.ID == id {
			indice = i
			flag = true
			src.ID = v.ID
			src.Name = v.Name
			src.Quantity = v.Quantity
			src.Code_value = v.Code_value
			src.Is_published = v.Is_published
			src.Expiration = v.Expiration
			src.Price = v.Price

		}
	}
	if flag == false {
		err = errors.New("Not Found")
		return
	}
	fmt.Println(src, src.ID)
	st := reflect.TypeOf(src)
	fmt.Println(st)
	field0 := st.Field(0)
	field1 := st.Field(1)
	field2 := st.Field(2)
	field3 := st.Field(3)
	field4 := st.Field(4)
	field5 := st.Field(5)
	field6 := st.Field(6)
	fmt.Println(string(field1.Tag.Get("val")))
	fmt.Println(string(field0.Tag.Get("val")))
	fmt.Println(string(field2.Tag.Get("val")))
	fmt.Println(string(field3.Tag.Get("val")))
	fmt.Println(string(field4.Tag.Get("val")))
	fmt.Println(string(field5.Tag.Get("val")))
	fmt.Println(string(field6.Tag.Get("val")))
	//vamos a parchar solos los atributos que estan
	for k, v := range data {
		switch {
		case k == string(field0.Tag.Get("val")):
			fmt.Println(k)

		case k == string(field1.Tag.Get("val")):
			//en esta instancia ya viene validado todo no usamos el ok
			val, _ := v.(string)
			src.Name = val

		case k == string(field2.Tag.Get("val")):
			//en esta instancia ya viene validado todo no usamos el ok
			val, _ := v.(int)
			fmt.Println(val)
			src.Quantity = int(val)

		case k == string(field3.Tag.Get("val")):
			//en esta instancia ya viene validado todo no usamos el ok
			val, _ := v.(string)
			src.Code_value = val

		case k == string(field4.Tag.Get("val")):
			// en esta instancia ya viene validado todo no usamos el ok
			val, _ := v.(bool)
			src.Is_published = val

		case k == string(field5.Tag.Get("val")):
			//en esta instancia ya viene validado todo no usamos el ok
			val, _ := v.(string)
			src.Expiration = val

		case k == string(field6.Tag.Get("val")):
			val, _ := v.(float64)
			src.Price = val
		default:
		}

	}
	//ya tenemos parchado el item
	fmt.Println(src)
	//cambiamos de struc a la bd
	dtapche := domain.Product{
		ID:           src.ID,
		Name:         src.Name,
		Quantity:     src.Quantity,
		Code_value:   src.Code_value,
		Is_published: src.Is_published,
		Expiration:   src.Expiration,
		Price:        src.Price,
	}

	//ahora agregamos el parche a la bd
	r.bd[indice] = dtapche

	ret = &dtapche
	return

}

func (r *RepositoryImplTesting) GetPriceMayor(price float64) (dt []domain.Product, err error) {
	flag := false

	var list []domain.Product
	for _, product := range r.bd {
		if product.Price > price {
			flag = true
			list = append(list, product)
		}
	}
	if flag == false {
		err = errors.New("Not Found")
		return
	}
	dt = list
	return

}

func (r *RepositoryImplTesting) GetById(id int) (ret *domain.Product, err error) {
	flag := false

	fmt.Println(id)
	for _, v := range r.bd {
		if v.ID == id {
			ret = &v
			//flag = true
			return
		}
	}
	if flag == false {
		err = errors.New("Item Not Found..")
		return
	}
	if flag {
		fmt.Println(ret)
	}

	return
}
func (r *RepositoryImplTesting) SaveNewProduct(data domain.Product) (ret *domain.Product, err error) {

	indi_candidate := len(r.bd) + 1
	for _, v := range r.bd {
		if indi_candidate == v.ID {
			indi_candidate++
		}
	}
	//despues ya tengo indice
	data.ID = indi_candidate
	r.bd = append(r.bd, data)

	return
}
