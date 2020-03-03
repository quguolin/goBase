package binding

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// cache struct reflect type bindCache.
var cache = &bindCache{
	data: make(map[reflect.Type]*sinfo),
}

var cacheTags = []string{"form", "header", "request"}

type bindCache struct {
	data  map[reflect.Type]*sinfo
	mutex sync.RWMutex
}

func (c *bindCache) get(obj reflect.Type) (s *sinfo) {
	var ok bool
	c.mutex.RLock()
	if s, ok = c.data[obj]; !ok {
		c.mutex.RUnlock()
		s = c.set(obj)
		return
	}
	c.mutex.RUnlock()
	return
}

func (c *bindCache) set(obj reflect.Type) (s *sinfo) {
	s = new(sinfo)
	tp := obj.Elem()
	for i := 0; i < tp.NumField(); i++ {
		fd := &field{
			tags:          make(map[string]tag, len(cacheTags)),
			defaultValues: make(map[string]defaultValue, len(cacheTags)),
		}
		fd.tp = tp.Field(i)
		for _, tagName := range cacheTags {
			rawTag := fd.tp.Tag.Get(tagName)
			if rawTag == "" {
				continue
			}
			parsedName, option := parseTag(rawTag)
			fd.tags[tagName] = tag{
				name:   parsedName,
				option: option,
			}
			if defV := fd.tp.Tag.Get("default"); defV != "" {
				dv := reflect.New(fd.tp.Type).Elem()
				setType(fd.tp.Type.Kind(), []string{defV}, dv, option)
				fd.defaultValues[tagName] = defaultValue{
					hasDefault: true,
					value:      dv,
				}
			}
		}

		s.field = append(s.field, fd)
	}
	c.mutex.Lock()
	c.data[obj] = s
	c.mutex.Unlock()
	return
}

type sinfo struct {
	field []*field
}

type field struct {
	tp            reflect.StructField
	tags          map[string]tag
	defaultValues map[string]defaultValue
}

type tag struct {
	name   string
	option tagOptions
}

type defaultValue struct {
	hasDefault bool
	value      reflect.Value
}

func (f *field) tag(tagName string) tag {
	return f.tags[tagName]
}

func (f *field) defaultValue(tagName string) defaultValue {
	return f.defaultValues[tagName]
}

func assign(ptr interface{}, form map[string][]string) error {
	sinfo := cache.get(reflect.TypeOf(ptr))
	val := reflect.ValueOf(ptr).Elem()
	for i, fd := range sinfo.field {
		typeField := fd.tp
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}

		tag := fd.tag("form")
		structFieldKind := structField.Kind()
		inputFieldName := tag.name
		if inputFieldName == "" {
			inputFieldName = typeField.Name

			// if "form" tag is nil, we inspect if the field is a struct.
			// this would not make sense for JSON parsing but it does for a form
			// since data is flatten
			if structFieldKind == reflect.Struct {
				err := assign(structField.Addr().Interface(), form)
				if err != nil {
					return err
				}
				continue
			}
		}

		dv := fd.defaultValue("form")
		inputValue, exists := form[inputFieldName]
		if !exists {
			// Set the field as default value when the input value is not exist
			if dv.hasDefault {
				structField.Set(dv.value)
			}
			continue
		}
		// Set the field as default value when the input value is empty
		if dv.hasDefault && inputValue[0] == "" {
			structField.Set(dv.value)
			continue
		}
		if _, isTime := structField.Interface().(time.Time); isTime {
			if err := setTimeField(inputValue[0], typeField, structField); err != nil {
				return err
			}
			continue
		}
		if err := setType(typeField.Type.Kind(), inputValue, structField, tag.option); err != nil {
			return err
		}
	}
	return nil
}

func setType(valueKind reflect.Kind, val []string, structField reflect.Value, option tagOptions) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val[0], 0, structField)
	case reflect.Int8:
		return setIntField(val[0], 8, structField)
	case reflect.Int16:
		return setIntField(val[0], 16, structField)
	case reflect.Int32:
		return setIntField(val[0], 32, structField)
	case reflect.Int64:
		return setIntField(val[0], 64, structField)
	case reflect.Uint:
		return setUintField(val[0], 0, structField)
	case reflect.Uint8:
		return setUintField(val[0], 8, structField)
	case reflect.Uint16:
		return setUintField(val[0], 16, structField)
	case reflect.Uint32:
		return setUintField(val[0], 32, structField)
	case reflect.Uint64:
		return setUintField(val[0], 64, structField)
	case reflect.Bool:
		return setBoolField(val[0], structField)
	case reflect.Float32:
		return setFloatField(val[0], 32, structField)
	case reflect.Float64:
		return setFloatField(val[0], 64, structField)
	case reflect.String:
		structField.SetString(val[0])
	case reflect.Slice:
		if option.Contains("split") {
			val = strings.Split(val[0], ",")
		}
		filtered := filterEmpty(val)
		sliceOf := structField.Type().Elem().Kind()
		numElems := len(filtered)
		slice := reflect.MakeSlice(structField.Type(), len(filtered), len(filtered))
		for i := 0; i < numElems; i++ {
			if err := setType(sliceOf, filtered[i:], slice.Index(i), ""); err != nil {
				return err
			}
		}
		structField.Set(slice)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	val = strings.TrimSpace(val)
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return errors.WithStack(err)
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	val = strings.TrimSpace(val)
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return errors.WithStack(err)
}

func setBoolField(val string, field reflect.Value) error {
	val = strings.TrimSpace(val)
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return nil
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	val = strings.TrimSpace(val)
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return errors.WithStack(err)
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	if timeFormat == "" {
		return errors.New("Blank time format")
	}
	val = strings.TrimSpace(val)
	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return errors.WithStack(err)
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return errors.WithStack(err)
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func filterEmpty(val []string) []string {
	filtered := make([]string, 0, len(val))
	for _, v := range val {
		v = strings.TrimSpace(v)
		if v != "" {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
