package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// Validate 输入验证的方法   数据验证
func Validate(data interface{}) (int, string) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans.New()) //转换成汉语，中文
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return 500, v.Translate(trans)
		}
	}
	return 200, ""
}
