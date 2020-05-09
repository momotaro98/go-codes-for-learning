package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type structTesg struct {
	// Should02 string `validate:"is-02"`
	GreaterThanOrEqualTo0 string `validate:"greater_than_or_equal_to_0"`
	// GreaterThan0OrEqualTo0 string `binding:"greater_than_or_equal_to_0"`
}

func main() {
	validate = validator.New()
	//validate.RegisterValidation("is-02", func(fl validator.FieldLevel) bool {
	validate.RegisterValidation("greater_than_or_equal_to_0", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "02"
	})

	s01 := &structTesg{
		GreaterThanOrEqualTo0: "01",
	}

	err := validate.Struct(s01)
	if err != nil {
		fmt.Printf("Err(s01):\n%+v\n", err)
	}

	s02 := &structTesg{
		GreaterThanOrEqualTo0: "02",
	}

	err = validate.Struct(s02)
	if err != nil {
		fmt.Printf("Err(s02):\n%+v\n", err)
	}
}
