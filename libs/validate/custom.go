package validate

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"regexp"
	"strings"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	translateOverride(trans)

	_ = validate.RegisterValidation("positive-number", PositiveNumber)
}

func PositiveNumber(fl validator.FieldLevel) bool {
	if _, ok := fl.Field().Interface().(int); ok {
		return true
	}
	return false

}

func renamejson(ss string) string {
	reg, _ := regexp.Compile("([A-Z])")
	sss := reg.ReplaceAllString(ss, "_${1}")
	return strings.Trim(strings.ToLower(sss), "_")
}

func translateOverride(trans ut.Translator) {

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}-不能为空", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", renamejson(fe.Field()))
		return t
	})

	_ = validate.RegisterTranslation("positive-number", trans, func(ut ut.Translator) error {
		return ut.Add("positive-number", "{0}-必须是个数字", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("positive-number", renamejson(fe.Field()))
		return t
	})

	_ = validate.RegisterTranslation("lt", trans, func(ut ut.Translator) error {
		return ut.Add("lt", "{0}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		if fe.Type().String() == "int" {
			t, _ := ut.T("lt", renamejson(fe.Field())+"-大小小于"+fe.Param())
			return t
		} else {
			t, _ := ut.T("lt", renamejson(fe.Field())+"-字符长度小于"+fe.Param())
			return t
		}
	})

	_ = validate.RegisterTranslation("gt", trans, func(ut ut.Translator) error {
		return ut.Add("gt", "{0}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		if fe.Type().String() == "int" {
			t, _ := ut.T("gt", renamejson(fe.Field())+"-大小大于"+fe.Param())
			return t
		} else {
			t, _ := ut.T("gt", renamejson(fe.Field())+"-字符长度大于"+fe.Param())
			return t
		}
	})

}
