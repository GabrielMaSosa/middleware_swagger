package domain

// con tag para reflect
type ProductPatch struct {
	Name         string  `Val_1:"name"`
	Quantity     int     `Val_2:"quantity"`
	Code_value   string  `Val_3:"code_value"`
	Is_published bool    `Val_4:"is_published"`
	Expiration   string  `Val_5:"expiration"`
	Price        float64 `Val_6:"price"`
}
