package validator

import (
	"go-crud/model/web"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CategoryValidateCreate(req web.CategoryCreateRequest) error {
	return validate.Struct(req)
}

func CategoryValidateUpdate(req web.CategoryUpdateRequest) error {
	return validate.Struct(req)
}
