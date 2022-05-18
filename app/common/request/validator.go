package request

import "github.com/go-playground/validator/v10"

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

//get error msg
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {
			//need request struct come true Validator interface
			if isValidator {
				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}

	return "Parameter error"
}
