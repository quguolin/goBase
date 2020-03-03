package binding

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *Validator) NewOnce() {
	v.once.Do(func() {
		v.validate = validator.New()
	})
}

func dataType(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func (v *Validator) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return errors.WithStack(err)
	}
	if err := assign(obj, req.Form); err != nil {
		return err
	}
	return v.ValidateStruct(obj)
}

func (v *Validator) ValidateStruct(obj interface{}) error {
	if dataType(obj) == reflect.Struct {
		v.NewOnce()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}
