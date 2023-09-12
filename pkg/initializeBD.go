package pkg

import (
	"encoding/json"
	"os"

	"github.com/GabrielMaSosa/middleware-swagger/internal/domain"
)

func InitilizeBD(path string) (dt []domain.Product, err error) {
	fil, err1 := os.ReadFile(path)
	if err1 != nil {
		err = err1
		return
	}
	json.Unmarshal(fil, &dt)

	return
}
