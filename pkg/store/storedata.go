package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
)

var (
	ErrReadDB  = errors.New("error read DB")
	ErrWRiteDB = errors.New("error Write DB")
)

func ReadAll(data string) (ret []domain.Product, err error) {
	fil, err1 := os.ReadFile(data)
	if err1 != nil {
		err = ErrReadDB
		return
	}
	json.Unmarshal(fil, &ret)

	return

}

func WriteAll(path string, in []domain.Product) (err error) {
	file, err2 := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err2 != nil {
		err = ErrWRiteDB
		return
	}
	json.NewEncoder(file).Encode(in)

	return

}
