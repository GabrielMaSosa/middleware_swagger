package domain

type Product struct {
	ID           int     `json:"id,omitempty,"`
	Name         string  `json:"name,omitempty,"`
	Quantity     int     `json:"quantity,omitempty,"`
	Code_value   string  `json:"code_value,omitempty,"`
	Is_published bool    `json:"is_published,omitempty,"`
	Expiration   string  `json:"expiration,omitempty,"`
	Price        float64 `json:"price,omitempty,"`
}
