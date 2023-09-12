package store

import (
	"encoding/json"
	"os"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
)

func ReadAll(data string) (ret []domain.Product, err error) {
	fil, err1 := os.ReadFile(data)
	if err1 != nil {
		err = err1
		return
	}
	json.Unmarshal(fil, &ret)

	return

}

func WriteAll(path string, in []domain.Product) (err error) {
	file, err2 := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err2 != nil {
		err = err2
		return
	}
	json.NewEncoder(file).Encode(in)

	return

}
