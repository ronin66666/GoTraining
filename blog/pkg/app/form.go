package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator" //通用翻译器，是一个使用 CLDR 数据 + 复数规则的 Go 语言 i18n 转换器。
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, validError := range v {
		errs = append(errs, validError.Error())
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// BindAndValid 针对入参校验的方法进行了二次封装
// 通过 ShouldBind 进行参数绑定和入参校验
// 当发生错误后，再通过上一步在中间件 Translations 设置的 Translator 来对错误消息体进行具体的翻译行为。
func BindAndValid(c *gin.Context, i interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(i)
	if err != nil {
		v := c.Value("trans") //获取中间件翻译器Translator
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
