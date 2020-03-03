package main

import (
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
)

// User contains user information
type User struct {
	FirstName string `validate:"required"`
	//LastName       string `validate:"required"`
	Age uint8 `validate:"gte=0,lte=130"`
	//Email          string `validate:"required,email"`
	//FavouriteColor string `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}

type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *Validator) NewOnce() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func (v *Validator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.NewOnce()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	user := &User{
		FirstName: "test",
		Age:       100,
	}
	val := &Validator{}
	err := val.ValidateStruct(user)
	if err != nil {
		panic(err)
	}
}
