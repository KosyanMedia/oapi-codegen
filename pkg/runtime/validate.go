package runtime

import (
	"encoding/base64"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

func init() {
	validate = validator.New()
	regexps = regexpCache{
		regexps: make(map[string]*regexp.Regexp),
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// TODO: Add xml support
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := validate.RegisterValidation("pattern", checkPattern)
	if err != nil {
		panic(err)
	}
}

var validate *validator.Validate
var regexps regexpCache

func ValidateInput(params interface{}) error {
	if params == nil {
		return nil
	}

	err := validate.Struct(params)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return validationErrors
	}

	return nil
}

func checkPattern(fl validator.FieldLevel) bool {
	param := fl.Param()
	if param == "" {
		return true
	}

	field := fl.Field()
	kind := field.Kind()

	switch kind {
	case reflect.String:
		regex, err := regexps.get(param)
		if err != nil {
			return false
		}
		value := field.String()
		return regex.MatchString(value)
	default:
		return false
	}
}

type regexpCache struct {
	mux     sync.RWMutex
	regexps map[string]*regexp.Regexp
}

func (c *regexpCache) get(param string) (*regexp.Regexp, error) {
	c.mux.RLock()
	if regex, ok := c.regexps[param]; ok {
		c.mux.RUnlock()
		return regex, nil
	}
	c.mux.RUnlock()

	pattern, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(fmt.Sprintf("^%s$", string(pattern)))
	if err != nil {
		return nil, err
	}

	c.mux.Lock()
	c.regexps[param] = regex
	c.mux.Unlock()

	return regex, nil
}
