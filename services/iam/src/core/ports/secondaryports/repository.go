package secondaryports

import (
	"iam/src/core/model"
	"iam/src/core/types"
)

type UsersRepository interface {
	SaveUser(user model.UserModel, authMethod model.UsersAuthMethodsModel) error
	GetUserById(id types.Id) (model.UserModel, error)
	ValidateEmail(userId types.Id) error
}

type EmailVerificationCodeRepository interface {
	SaveCode(code model.EmailVerificationCodeModel) error
	CountActiveCodes(userId types.Id) (int, error)
	GetCode(code types.Code) (model.EmailVerificationCodeModel, error)
	DeleteCode(code types.Code) error
}

type AuthorizationCodeRepository interface {
	SaveCode(code model.AuthorizationCodeModel) error
	// CountActiveCodes(userId types.Id) (int, error)
	GetCode(code types.Code) (model.AuthorizationCodeModel, error)
	DeleteCode(code types.Code) error
}
