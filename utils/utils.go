package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"sync"
)

var validate *validator.Validate
var validateOnce sync.Once

func getValidate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}

func CheckParams(requestBody string, param interface{}) error {
	var err error
	if err = json.Unmarshal([]byte(requestBody), param); err != nil {
		return err
	}
	if err = getValidate().Struct(param); err != nil {
		return err
	}
	return nil
}

